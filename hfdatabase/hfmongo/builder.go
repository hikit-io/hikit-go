package hfmongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Builder struct {
	fields        map[string]*Filed
	FindOptions   *options.FindOptions
	UpdateOptions *options.UpdateOptions
}

func (o *Builder) Reset() {
	o.init()
}

func (o *Builder) init() {
	if o.fields == nil {
		o.fields = map[string]*Filed{}
	}
	if o.FindOptions == nil {
		o.FindOptions = options.Find()
	}
	if o.UpdateOptions == nil {
		o.UpdateOptions = options.Update()
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
