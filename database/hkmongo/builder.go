package hkmongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	builderPool = sync.Pool{
		New: func() interface{} {
			b := &Builder{}
			b.init()
			return b
		},
	}
)

func NewBuilder() *Builder {
	return builderPool.Get().(*Builder).Reset()
}

//Free 之后不可再用
func (b *Builder) Free() {
	builderPool.Put(b)
}

type Builder struct {
	fields        map[string]*Field
	logicFields   map[string][]bson.D
	findOptions   *options.FindOptions
	updateOptions *options.UpdateOptions
}

func (b *Builder) Reset() *Builder {
	b.fields = map[string]*Field{}
	b.findOptions = &options.FindOptions{}
	b.updateOptions = &options.UpdateOptions{}
	b.logicFields = map[string][]bson.D{}
	return b
}

func (b *Builder) Skip(count int64) *Builder {
	b.init()
	b.findOptions.Skip = &count
	return b
}

func (b *Builder) Limit(count int64) *Builder {
	b.init()
	b.findOptions.Limit = &count
	return b
}

func (b *Builder) Upsert(enable bool) *Builder {
	b.init()
	b.updateOptions.SetUpsert(enable)
	return b
}

func (b *Builder) BypassDocumentValidation(enable bool) *Builder {
	b.init()
	b.updateOptions.SetBypassDocumentValidation(enable)
	return b
}

func (b *Builder) init() {
	if b.fields == nil {
		b.fields = map[string]*Field{}
	}
	if b.findOptions == nil {
		b.findOptions = options.Find()
	}
	if b.updateOptions == nil {
		b.updateOptions = options.Update()
	}
	if b.logicFields == nil {
		b.logicFields = map[string][]bson.D{}
	}
}

func (b *Builder) Field(name string) *Field {
	b.init()
	if b.fields[name] == nil {
		field := &Field{
			map[FieldName]*Field{},
			name,
			D{},
		}
		b.fields[name] = field
	}
	return b.fields[name]
}

func mergeFindField(prefix string, fields map[string]*Field) bson.D {
	all := bson.D{}
	for _, filed := range fields {
		bsonD := bson.D{}
		for _, e := range filed.val {
			if e.opType == OpTypeFind {
				bsonD = append(bsonD, bson.E{
					Key:   e.Key,
					Value: e.Value,
				})
			}
		}

		name := filed.name
		if prefix != "" {
			name = fmt.Sprintf("%s.%s", prefix, filed.name)
		}

		if filed.chs != nil {
			bd := mergeFindField(name, filed.chs)
			all = append(all, bd...)
		}

		if len(bsonD) != 0 {

			bsonE := bson.E{
				Key:   name,
				Value: bsonD,
			}
			all = append(all, bsonE)
		}
	}
	return all
}

func (b *Builder) Filter() bson.D {
	b.init()
	all := mergeFindField("", b.fields)
	for opName, bd := range b.logicFields {
		if len(bd) != 0 {
			all = append(all, bson.E{
				Key:   opName,
				Value: b.logicFields[opName],
			})
		}
	}
	return all
	//for _, filed := range b.fields {
	//	bsonD := bson.D{}
	//	for _, e := range filed.val {
	//		if e.opType == OpTypeFind {
	//			bsonD = append(bsonD, bson.E{
	//				Key:   e.Key,
	//				Value: e.Value,
	//			})
	//		}
	//	}
	//	if len(bsonD) != 0 {
	//		bsonE := bson.E{
	//			Key:   filed.name,
	//			Value: bsonD,
	//		}
	//		all = append(all, bsonE)
	//	}
	//}
	//return all
}

func mergeUpField(prefix string, fields map[string]*Field) map[OpName]bson.D {
	opMap := map[OpName]bson.D{}

	for _, filed := range fields {
		name := filed.name
		if prefix != "" {
			name = fmt.Sprintf("%s.%s", prefix, filed.name)
		}
		for _, e := range filed.val {
			if e.opType == OpTypeUpdate {
				opMap[e.Key] = append(opMap[e.Key], bson.E{
					Key:   name,
					Value: e.Value,
				})
			}
		}

		if filed.chs != nil {
			bd := mergeUpField(name, filed.chs)
			for opName, d := range bd {
				opMap[opName] = append(opMap[opName], d...)
			}
		}
	}

	return opMap
}

func (b *Builder) Update() bson.D {
	b.init()
	all := bson.D{}
	m := mergeUpField("", b.fields)
	for opName, d := range m {
		all = append(all, bson.E{
			Key:   opName,
			Value: d,
		})
	}
	return all
	//for _, filed := range b.fields {
	//	opMap := map[OpName]bson.D{}
	//	for _, e := range filed.val {
	//		if e.opType == OpTypeUpdate {
	//			opMap[e.Key] = append(opMap[e.Key], bson.E{
	//				Key:   filed.name,
	//				Value: e.Value,
	//			})
	//		}
	//	}
	//	for op, d := range opMap {
	//		bsonE := bson.E{
	//			Key:   op,
	//			Value: d,
	//		}
	//		all = append(all, bsonE)
	//	}
	//}
	//return all
}

func (b *Builder) FindOpts() *options.FindOptions {
	b.init()
	var (
		sort bson.D
		pros bson.D
		hint bson.D
	)

	for _, filed := range b.fields {
		for _, e := range filed.val {
			if e.opType == OpTypeQuery {
				switch e.Key {
				case QueryOp.Sort:
					sort = append(sort, bson.E{
						Key:   filed.name,
						Value: e.Value,
					})
				case QueryOp.Projection:
					pros = append(pros, bson.E{
						Key:   filed.name,
						Value: e.Value,
					})
				case QueryOp.Hint:
					hint = append(hint, bson.E{
						Key:   filed.name,
						Value: e.Value,
					})
				}
			}
		}

	}
	if len(sort) != 0 {
		b.findOptions.Sort = sort
	}
	if len(pros) != 0 {
		b.findOptions.Projection = pros
	}
	if len(hint) != 0 {
		b.findOptions.Hint = hint
	}
	return b.findOptions
}

func (b *Builder) UpOpts() *options.UpdateOptions {
	b.init()
	var (
		hint bson.D
	)
	for _, filed := range b.fields {
		for _, e := range filed.val {
			if e.opType == OpTypeQuery {
				switch e.Key {
				case QueryOp.Hint:
					hint = append(hint, bson.E{
						Key:   filed.name,
						Value: e.Value,
					})
				}
			}
		}

	}

	if len(hint) != 0 {
		b.updateOptions.Hint = hint
	}
	return b.updateOptions
}

type BuilderOrFc = interface{}

func (b *Builder) Or(bs *Builder) *Builder {
	b.logicFields[LogicOp.Or] = append(b.logicFields[LogicOp.Or], bs.Filter())
	return b
}

func (b *Builder) OrFc(fc func(br *Builder)) *Builder {
	bs := NewBuilder()
	fc(bs)
	b.logicFields[LogicOp.Or] = append(b.logicFields[LogicOp.Or], bs.Filter())
	bs.Free()
	return b
}

func (b *Builder) Nor(bs *Builder) *Builder {
	b.logicFields[LogicOp.Nor] = append(b.logicFields[LogicOp.Nor], bs.Filter())
	return b
}

func (b *Builder) NorFc(fc func(br *Builder)) *Builder {
	bs := NewBuilder()
	fc(bs)
	b.logicFields[LogicOp.Nor] = append(b.logicFields[LogicOp.Nor], bs.Filter())
	bs.Free()
	return b
}
