package hftypes

import (
	"errors"
	"fmt"
	"reflect"
)

type (
	MapStrI   = map[Str]I
	MapStrI8  = map[Str]I8
	MapStrI16 = map[Str]I16
	MapStrI32 = map[Str]I32
	MapStrI64 = map[Str]I64

	MapStrUi   = map[Str]Ui
	MapStrUi8  = map[Str]Ui8
	MapStrUi16 = map[Str]Ui16
	MapStrUi32 = map[Str]Ui32
	MapStrUi64 = map[Str]Ui64

	MapStrF32 = map[Str]F32
	MapStrF64 = map[Str]F64

	MapStrStr = map[Str]Str
	MapStrAny = map[Str]Any
)

func MapStrAnyToStruct(m MustMapStrAny, s MustStructPtr) error {
	var (
		mv reflect.Value
	)
	switch {
	case IsMapStrAny(m):
		mv = reflect.ValueOf(m)
	case IsMapStrAnyPtr(m):
		mv = reflect.ValueOf(m).Elem()
	}
	if !IsStructPtr(s) {
		return errors.New("s must be struct ptr")
	}
	sv := reflect.ValueOf(s)
	mi := mv.MapRange()
	for mi.Next() {
		field := sv.Elem().FieldByName(mi.Key().String())
		if field.Kind() == mi.Value().Kind() {
			field.Set(mi.Value())
		} else {
			if field.Kind() == mi.Value().Elem().Kind() {
				field.Set(mi.Value().Elem())
				continue
			}
			if mi.Value().Elem().Kind() == reflect.Map {
				err := mapToStruct(mi.Value().Elem(), field)
				if err != nil {
					return err
				}
				continue
			}
			return errors.New(fmt.Sprintf("map [%v] kind is %v,struct key kind is %v", mi.Key().String(), mi.Value().Kind(), field.Kind()))
		}
	}
	return nil
}

func mapToStruct(m reflect.Value, s reflect.Value) error {
	mv := m
	sv := s
	mi := mv.MapRange()
	for mi.Next() {
		field := sv.FieldByName(mi.Key().String())
		if field.Kind() == mi.Value().Kind() {
			field.Set(mi.Value())
		} else {
			if field.Kind() == mi.Value().Elem().Kind() {
				field.Set(mi.Value().Elem())
				continue
			}
			if mi.Value().Kind() == reflect.Map {
				err := mapToStruct(mi.Value().Elem(), field)
				if err != nil {
					return err
				}
				continue
			}
			return errors.New(fmt.Sprintf("map [%v] kind is %v,struct key kind is %v", mi.Key().String(), mi.Value().Kind(), field.Kind()))
		}
	}
	return nil
}
