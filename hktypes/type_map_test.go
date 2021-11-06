package hktypes

import (
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	m := map[string]interface{}{
		"Key":  "val",
		"Key1": "2",
		"Key2": map[string]interface{}{
			"Key3": "4",
		},
	}
	s := struct {
		Key  string
		Key1 string
		Key2 struct {
			Key3 string
		}
	}{
		Key:  "1",
		Key1: "3",
		Key2: struct{ Key3 string }{Key3: "54"},
	}
	err := MapStrAnyToStruct(m, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
