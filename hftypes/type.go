package hftypes

import (
	"reflect"
	"strconv"
)

const (
	IntSize = strconv.IntSize
	Bits64  = 64
	Bits32  = 32
)

const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

type (
	Any  = interface{}
	Anys = []interface{}
	Str  = string
	B    = bool
	F32  = float32
	F64  = float64
	I8   = int8
	I16  = int16
	I32  = int32
	I64  = int64
	I    = int
	Ui8  = uint8
	Ui16 = uint16
	Ui32 = uint32
	Ui64 = uint64
	Ui   = uint

	MustPtr = Any

	MustStruct      = Any
	MustStructPtr   = Any
	MustStructOrPtr = Any

	MustSlice      = Any
	MustSlicePtr   = Any
	MustSliceOrPtr = Any

	MustSliceInt      = Any
	MustSliceIntPtr   = Any
	MustSliceIntOrPtr = Any

	MustSliceFloat      = Any
	MustSliceFloatPtr   = Any
	MustSliceFloatOrPtr = Any

	MustSliceStr      = Any
	MustSliceStrPtr   = Any
	MustSliceStrOrPtr = Any

	MustSliceStruct      = Any
	MustSliceStructPtr   = Any
	MustSliceStructOrPtr = Any

	MustInt      = Any
	MustIntPtr   = Any
	MustIntOrPtr = Any

	MustFloat      = Any
	MustFloatPtr   = Any
	MustFloatOrPtr = Any

	MustStr      = Any
	MustStrPtr   = Any
	MustStrOrPtr = Any

	MustMap            = Any
	MustMapPtr         = Any
	MustMapOrPtr       = Any
	MustMapStrAny      = Any
	MustMapStrAnyPtr   = Any
	MustMapStrAnyOrPtr = Any
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

func IsMapStrAny(m Any) B {
	if m == nil {
		return false
	}
	_, ok := m.(map[string]Any)
	return ok
}

func IsMapStrAnyPtr(m Any) B {
	if m == nil {
		return false
	}
	_, ok := m.(*map[string]Any)
	return ok
}

func IsMapStrAnyOrPtr(m Any) B {
	if IsMapStrAny(m) {
		return true
	}
	return IsMapStrAnyPtr(m)
}

func IsStructOrPtr(e Any) B {
	v := reflect.TypeOf(e)
	if v.Kind() == reflect.Struct {
		return true
	}
	if v.Kind() == reflect.Ptr {
		if v.Elem().Kind() == reflect.Struct {
			return true
		}
	}
	return false
}

func IsStruct(e Any) B {
	if e == nil {
		return false
	}
	v := reflect.TypeOf(e)
	if v.Kind() == reflect.Struct {
		return true
	}
	return false
}

func IsStructPtr(e Any) B {
	if e == nil {
		return false
	}
	v := reflect.TypeOf(e)
	if v.Kind() == reflect.Ptr {
		if v.Elem().Kind() == reflect.Struct {
			return true
		}
	}
	return false
}

type TypeFc func(elem Any) bool

func IsAny(elem Any) B {
	_, ok := elem.(interface{})
	return ok
}

func IsInterface(elem Any) B {
	switch elem.(type) {
	case interface{}:
		return true
	}
	return false
}

func IsSliceOrPtr(e Any) B {
	if IsSlice(e) {
		return true
	}
	return IsSlicePtr(e)
}

func IsSlice(elem Any) B {
	// 已知类型断言
	switch elem.(type) {
	case []Any:
		return true
	case []Str:
		return true
	case []F32, []F64:
		return true
	case []Ui, []Ui8, []Ui16, []Ui32, []Ui64:
		return true
	case []I, []I8, []I16, []I32, []I64:
		return true
	}

	tv := reflect.ValueOf(elem)
	return tv.Kind() == reflect.Slice
}

func IsSlicePtr(elem Any) B {
	// 已知类型断言
	switch elem.(type) {
	case []Any:
		return true
	case []*Str:
		return true
	case []*F32, []*F64:
		return true
	case []*Ui, []*Ui8, []*Ui16, []*Ui32, []*Ui64:
		return true
	case []*I, []*I8, []*I16, []*I32, []*I64:
		return true
	}

	tv := reflect.ValueOf(elem)
	if tv.Kind() == reflect.Ptr {
		return tv.Elem().Kind() == reflect.Slice
	}
	return false
}

func IsAnys(elem Any) B {
	_, ok := elem.([]Any)
	return ok
}
