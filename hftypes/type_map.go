package hftypes

import "reflect"

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
)

func IsMap(e Any) B {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Map {
		return true
	}
	return false
}

func IsMapPtr(e Any) B {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() == reflect.Map {
			return true
		}
	}
	return false
}

func IsMapOrPtr(e Any) B {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Map {
		return true
	}
	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() == reflect.Map {
			return true
		}
	}
	return false
}
