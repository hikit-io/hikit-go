package hftypes

import (
	"fmt"
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
)

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

func IsSlice(elem Any) B {
	// 已知类型断言
	switch elem.(type) {
	case []Any, *[]Any, []I, *[]I, []Str, *[]Str:
		return true
	}

	tv := reflect.ValueOf(elem)
	if tv.Kind() == reflect.Ptr {
		fmt.Println(tv.Elem().Kind())
		return tv.Elem().Kind() == reflect.Slice
	}
	return tv.Kind() == reflect.Slice
}

func IsAnys(elem Any) B {
	_, ok := elem.([]interface{})
	return ok
}

func Complex64(elem Any) B {
	_, ok := elem.(complex64)
	return ok
}

func Complex128(elem Any) B {
	_, ok := elem.(complex128)
	return ok
}
