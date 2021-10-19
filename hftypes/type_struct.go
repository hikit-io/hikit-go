package hftypes

import (
	"errors"
	"reflect"
)

func StructToMapStrStr(split Str, s MustStruct, m MustMapStrAny) {

}

func StructToMapStrAny(s MustStruct, m MustMapStrAnyOrPtr) error {
	var (
		mi map[Str]Any
		sv reflect.Value
	)
	switch {
	case IsMapStrAny(m):
		mi = m.(MapStrAny)
	case IsMapStrAnyPtr(m):
		mi = *(m.(*MapStrAny))
	default:
		return errors.New("s must be map[Str]Any or *map[Str]Any ")
	}

	switch {
	case IsStruct(s):
		sv = reflect.ValueOf(s)
	case IsStructPtr(s):
		sv = reflect.ValueOf(s).Elem()
	default:
		return errors.New("s must be struct or *struct")
	}

	for i := 0; i < sv.NumField(); i++ {
		field := sv.Field(i)
		if field.Kind() == reflect.Struct {
			im := map[Str]Any{}
			err := StructToMapStrAny(field.Interface(), im)
			if err != nil {
				return err
			}
			mi[sv.Type().Field(i).Name] = im
			continue
		}
		if field.Kind() == reflect.Ptr {
			if field.Elem().Kind() == reflect.Struct {
				im := map[Str]Any{}
				err := StructToMapStrAny(field.Elem().Interface(), im)
				if err != nil {
					return err
				}
				mi[sv.Type().Field(i).Name] = im
				continue
			}
			if !field.IsZero() {
				mi[sv.Type().Field(i).Name] = field.Elem().Interface()
				continue
			}
			if !field.IsNil() {
				mi[sv.Type().Field(i).Name] = field.Elem().Interface()
				continue
			}
		}
		mi[sv.Type().Field(i).Name] = field.Interface()
	}
	return nil
}
