package hftypes

type (
	Any  = interface{}
	Void = struct{}
	Str  = string
	B    = bool
	F32  = float32
	F64  = float64
	I8   = int8
	I    = int
	I16  = int16
	I32  = int32
	I64  = int64
	Ui8  = uint8
	Ui   = uint
	Ui16 = uint16
	Ui32 = uint32
	Ui64 = uint64
)

type TypeFc func(elem Any) bool

func IsAny(elem Any) B {
	_, ok := elem.(interface{})
	return ok
}

func IsStr(elem Any) B {
	_, ok := elem.(string)
	return ok
}

func IsB(elem Any) B {
	_, ok := elem.(bool)
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

func IsUi(elem Any) B {
	_, ok := elem.(uint)
	return ok
}

func IsUi8(elem Any) B {
	_, ok := elem.(uint8)
	return ok
}

func IsUi16(elem Any) B {
	_, ok := elem.(uint16)
	return ok
}

func IsUi32(elem Any) B {
	_, ok := elem.(uint32)
	return ok
}

func IsUi64(elem Any) B {
	_, ok := elem.(uint64)
	return ok
}

func IsI(elem Any) B {
	_, ok := elem.(int)
	return ok
}

func IsI8(elem Any) B {
	_, ok := elem.(int8)
	return ok
}

func IsI16(elem Any) B {
	_, ok := elem.(int16)
	return ok
}

func IsI32(elem Any) B {
	_, ok := elem.(int32)
	return ok
}

func IsI64(elem Any) B {
	_, ok := elem.(int64)
	return ok
}

func IsF32(elem Any) B {
	_, ok := elem.(float32)
	return ok
}

func IsF64(elem Any) B {
	_, ok := elem.(float64)
	return ok
}

func ToBool(e Any) B {
	return e.(bool)
}

func ToStr(e Any) string {
	return e.(string)
}

func ToI(e Any) I {
	return e.(int)
}

func ToI8(e Any) I8 {
	return e.(int8)
}

func ToI16(e Any) I16 {
	return e.(int16)
}

func ToI32(e Any) I32 {
	return e.(int32)
}

func ToI64(e Any) I64 {
	return e.(int64)
}

func ToUi(e Any) Ui {
	return e.(uint)
}

func ToUi8(e Any) Ui8 {
	return e.(uint8)
}

func ToUi16(e Any) Ui16 {
	return e.(uint16)
}

func ToUi32(e Any) Ui32 {
	return e.(uint32)
}

func ToUi64(e Any) Ui64 {
	return e.(uint64)
}

func AsBool(e Any) B {
	return e.(bool)
}

func AsStr(e Any) string {
	return e.(string)
}

func AsI(e Any) I {
	return e.(int)
}

func AsI8(e Any) I8 {
	return e.(I8)
}

func AsI16(e Any) I16 {
	return e.(int16)
}

func AsI32(e Any) I32 {
	return e.(int32)
}

func AsI64(e Any) I64 {
	return e.(int64)
}

func AsUi(e Any) Ui {
	return e.(uint)
}

func AsUi8(e Any) Ui8 {
	return e.(uint8)
}

func AsUi16(e Any) Ui16 {
	return e.(uint16)
}

func AsUi32(e Any) Ui32 {
	return e.(uint32)
}

func AsUi64(e Any) Ui64 {
	return e.(uint64)
}
