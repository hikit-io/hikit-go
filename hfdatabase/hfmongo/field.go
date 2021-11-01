package hfmongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
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
	childs map[FieldName]*Field
	name   FieldName
	val    D
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
	if f.childs == nil {
		f.childs = map[FieldName]*Field{}
	}
	if f.childs[name] == nil {
		f.childs[name] = &Field{
			name: name,
		}
	}
	return f.childs[name]
}
