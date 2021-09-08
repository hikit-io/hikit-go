package mongo

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
)

type E struct {
	bson.E
	opType OpType
}

type D []E

type Filed struct {
	name FieldName
	val  D
}

func (f *Filed) E() bson.E {
	return bson.E{
		Key:   f.name,
		Value: f.val,
	}
}

func (f *Filed) op(opName OpName, opType OpType, val Any) *Filed {
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
