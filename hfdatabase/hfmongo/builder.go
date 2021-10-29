package hfmongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Builder struct {
	fields        map[string]*Filed
	findOptions   *options.FindOptions
	updateOptions *options.UpdateOptions
}

func (o *Builder) Reset() {
	o.fields = map[string]*Filed{}
	o.findOptions = &options.FindOptions{}
	o.updateOptions = &options.UpdateOptions{}
}

func (o *Builder) Skip(count int64) {
	o.init()
	o.findOptions.Skip = &count
}

func (o *Builder) Limit(count int64) {
	o.init()
	o.findOptions.Limit = &count
}

func (o *Builder) init() {
	if o.fields == nil {
		o.fields = map[string]*Filed{}
	}
	if o.findOptions == nil {
		o.findOptions = options.Find()
	}
	if o.updateOptions == nil {
		o.updateOptions = options.Update()
	}
}

func (o *Builder) Field(name string) *Filed {
	o.init()
	if o.fields[name] == nil {
		field := &Filed{
			name,
			D{},
		}
		o.fields[name] = field
	}
	return o.fields[name]
}

func (o *Builder) Filter() bson.D {
	o.init()
	all := bson.D{}
	for _, filed := range o.fields {
		bsonD := bson.D{}
		for _, e := range filed.val {
			if e.opType == OpTypeFind {
				bsonD = append(bsonD, bson.E{
					Key:   e.Key,
					Value: e.Value,
				})
			}
		}
		if len(bsonD) != 0 {
			bsonE := bson.E{
				Key:   filed.name,
				Value: bsonD,
			}
			all = append(all, bsonE)
		}
	}
	return all
}

func (o *Builder) Update() bson.D {
	o.init()
	all := bson.D{}
	for _, filed := range o.fields {
		opMap := map[OpName]bson.D{}
		for _, e := range filed.val {
			if e.opType == OpTypeUpdate {
				opMap[e.Key] = append(opMap[e.Key], bson.E{
					Key:   filed.name,
					Value: e.Value,
				})
			}
		}
		for op, d := range opMap {
			bsonE := bson.E{
				Key:   op,
				Value: d,
			}
			all = append(all, bsonE)
		}
	}
	return all
}

func (o *Builder) Options() *options.FindOptions {
	o.init()
	var (
		sort bson.D
		pros bson.D
		hint bson.D
	)

	for _, filed := range o.fields {
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
					pros = append(pros, bson.E{
						Key:   filed.name,
						Value: e.Value,
					})
				}
			}
		}

	}
	if len(sort) != 0 {
		o.findOptions.Sort = sort
	}
	if len(pros) != 0 {
		o.findOptions.Projection = pros
	}
	if len(hint) != 0 {
		o.findOptions.Hint = hint
	}
	return o.findOptions
}
