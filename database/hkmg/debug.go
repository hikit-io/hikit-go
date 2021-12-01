package hkmg

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type DebugD bson.D

func (d DebugD) MarshalJSON() ([]byte, error) {
	return json.Marshal(bsonDToMap(bson.D(d)))
}

func bsonDToMap(d bson.D) map[string]interface{} {
	m := d.Map()
	nm := map[string]interface{}{}
	for k, e := range m {
		if v, ok := e.([]bson.D); ok {
			var s []map[string]interface{}
			for _, b := range v {
				s = append(s, bsonDToMap(b))
			}
			nm[k] = s
			continue
		}
		if v, ok := e.(bson.D); ok {
			nm[k] = bsonDToMap(v)
			continue
		}
		nm[k] = e
	}
	return nm
}
