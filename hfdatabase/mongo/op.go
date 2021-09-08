package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Op struct {
	find map[string]*Filed
}

func (o *Op) init() {
	if o.find == nil {
		o.find = map[string]*Filed{}
	}

}

func (o *Op) Field(name string) *Filed {
	o.init()
	if o.find[name] == nil {
		field := &Filed{
			name,
			D{},
		}
		o.find[name] = field
	}
	return o.find[name]
}

func (o *Op) Find() bson.D {
	o.init()
	all := bson.D{}
	for _, filed := range o.find {
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

func (o *Op) Update() bson.D {
	o.init()
	all := bson.D{}
	for _, filed := range o.find {
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
