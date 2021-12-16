package hkmg

import (
	"go.mongodb.org/mongo-driver/bson"

	. "go.hikit.io/hktypes"
)

type FieldName = Str
type OpName = Str

type OpType uint8

const (
	OpTypeFind OpType = iota + 1
	OpTypeUpdate
	OpTypeQuery
)

type E struct {
	bson.E
	opType OpType
}

type D []E

type Field struct {
	chs  []*Field
	name FieldName
	val  D
}

func (f *Field) E() bson.E {
	return bson.E{
		Key:   f.name,
		Value: f.val,
	}
}

func (f *Field) op(opName OpName, opType OpType, val Any) *Field {
	for i := range f.val {
		if f.val[i].Key == opName {
			f.val[i].Value = val
			return f
		}
	}

	f.val = append(f.val, E{
		E: bson.E{
			Key:   opName,
			Value: val,
		},
		opType: opType,
	})
	return f
}

func (f *Field) Child(name string) *Field {
	if f.chs == nil {
		f.chs = []*Field{}
	}
	for _i, ch := range f.chs {
		if ch.name == name {
			return f.chs[_i]
		}
	}
	field := &Field{
		name: name,
	}
	//if f.chs[name] == nil {
	//	field = &Field{
	//		name: name,
	//	}
	//}
	f.chs = append(f.chs, field)
	return field
}

func In(val MustSlice) bson.E {
	return bson.E{
		Key:   FindOp.In,
		Value: val,
	}
}

func NotIn(val MustSlice) bson.E {
	return bson.E{
		Key:   FindOp.NotIn,
		Value: val,
	}
}

func LessThan(val Any) bson.E {
	return bson.E{
		Key:   FindOp.LessThan,
		Value: val,
	}
}

func LessThanEqual(val Any) bson.E {
	return bson.E{
		Key:   FindOp.LessThanEqual,
		Value: val,
	}
}

func Equal(val Any) bson.E {
	return bson.E{
		Key:   FindOp.Equal,
		Value: val,
	}
}

func GreatThan(val Any) bson.E {
	return bson.E{
		Key:   FindOp.GreatThan,
		Value: val,
	}
}

func GreatThanEqual(val Any) bson.E {
	return bson.E{
		Key:   FindOp.GreatThanEqual,
		Value: val,
	}
}

func All(val MustSlice) bson.E {
	return bson.E{
		Key:   FindOp.All,
		Value: val,
	}
}

func Regex(val string) bson.E {
	return bson.E{
		Key:   FindOp.Regex,
		Value: val,
	}
}
