package hftypes

import (
	"errors"
	"reflect"
)

func StructToMapStrStr() {

}

func StructToMapStrAny(s MustStruct, m MustMapStrAny) error {
	var (
		mi map[Str]Any
	)
	if IsMapStrAny(m) {
		mi = m.(MapStrAny)
	}
	if IsMapStrAnyPtr(m) {
		mi = *(m.(*MapStrAny))
	}
	if !IsStruct(s) {
		return errors.New("")
	}
	sv := reflect.ValueOf(s)
	for i := 0; i < sv.NumField(); i++ {
		if sv.Field(i).Kind() == reflect.Struct {
			im := map[Str]Any{}
			err := StructToMapStrAny(sv.Field(i).Interface(), im)
			if err != nil {
				return err
			}
			mi[sv.Type().Field(i).Name] = im
			continue
		}
		mi[sv.Type().Field(i).Name] = sv.Field(i).Interface()
	}
	return nil
}
