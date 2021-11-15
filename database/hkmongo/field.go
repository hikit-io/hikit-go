package hkmongo

import (
	. "go.hikit.io/hktypes"
	"go.mongodb.org/mongo-driver/bson"
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
	chs  map[FieldName]*Field
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
		f.chs = map[FieldName]*Field{}
	}
	if f.chs[name] == nil {
		f.chs[name] = &Field{
			name: name,
		}
	}
	return f.chs[name]
}
