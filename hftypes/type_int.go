package hftypes

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

func IsInt(elem Any) B {
	switch elem.(type) {
	case I, I8, I16, I32, I64, *I, *I8, *I16, *I32, *I64:
		return true
	}
	return false
}

func IsUint(elem Any) B {
	switch elem.(type) {
	case Ui, Ui8, Ui16, Ui32, Ui64, *Ui, *Ui8, *Ui16, *Ui32, *Ui64:
		return true
	}
	return false
}

func IsIs(elem Any) B {
	_, ok := elem.([]int)
	return ok
}

func IsI8s(elem Any) B {
	_, ok := elem.([]I8)
	return ok
}

func IsI16s(elem Any) B {
	_, ok := elem.([]I16)
	return ok
}

func IsI32s(elem Any) B {
	_, ok := elem.([]I32)
	return ok
}

func IsI64s(elem Any) B {
	_, ok := elem.([]I64)
	return ok
}

func IsUis(elem Any) B {
	_, ok := elem.([]Ui)
	return ok
}

func IsUi8s(elem Any) B {
	_, ok := elem.([]Ui8)
	return ok
}

func IsUi6s(elem Any) B {
	_, ok := elem.([]Ui16)
	return ok
}

func IsUi32s(elem Any) B {
	_, ok := elem.([]Ui32)
	return ok
}

func IsUi64s(elem Any) B {
	_, ok := elem.([]Ui64)
	return ok
}

func IsF32s(elem Any) B {
	_, ok := elem.([]F32)
	return ok
}

func IsF64s(elem Any) B {
	_, ok := elem.([]F64)
	return ok
}

func RangeI8(i64 I64) bool {
	return i64 >= MinInt8 && i64 <= MaxInt8
}

func RangeI16(i64 I64) bool {
	return i64 >= MinInt16 && i64 <= MaxInt16
}

func RangeI32(i64 I64) bool {
	return i64 >= MinInt32 && i64 <= MaxInt32
}

func AsI8(e Any) I8 {
	switch {
	case IsI8(e):
		return e.(I8)
	case IsUi(e) || IsUi8(e) || IsUi16(e) || IsUi32(e) || IsUi64(e):
		i := AsUi64(e)
		if RangeI8(I64(i)) {
			return I8(i)
		}
		panic(i)
	case IsI(e) || IsI16(e) || IsI32(e) || IsI64(e):
		i64 := AsI64(e)
		if RangeI8(i64) {
			return I8(i64)
		}
		panic(i64)
	default:
		panic("not support")
	}
}

func AsI16(e Any) I16 {
	switch {
	case IsI16(e):
		return e.(I16)
	case IsI8(e):
		return I16(AsI8(e))
	case IsUi8(e):
		return I16(AsUi8(e))
	case IsUi(e) || IsUi32(e) || IsUi64(e) || IsUi64(e):
		i := AsUi64(e)
		if RangeI16(I64(i)) {
			return I16(i)
		}
		panic(i)
	case IsI(e) || IsI32(e) || IsI64(e):
		i64 := AsI64(e)
		if RangeI16(i64) {
			return I16(i64)
		}
		panic(i64)
	default:
		panic("not support")
	}
}

func AsI32(e Any) I32 {
	switch {
	case IsUi(e) || IsUi32(e) || IsUi64(e) || IsUi64(e):
		i := AsUi64(e)
		if RangeI32(I64(i)) {
			return I32(i)
		}
		panic(i)
	case IsI(e) || IsI32(e) || IsI64(e) || IsI8(e):
		i64 := AsI64(e)
		if RangeI32(i64) {
			return I32(i64)
		}
		panic(i64)
	default:
		panic("not support")
	}
}

func AsI64(e Any) I64 {
	switch {
	case IsI(e):
		return I64(AsI(e))
	case IsI8(e):
		return I64(AsI8(e))
	case IsI16(e):
		return I64(AsI16(e))
	case IsI32(e):
		return I64(AsI32(e))
	case IsI64(e):
		return e.(I64)
	}
	panic("not support type")
}

func AsUi(e Any) Ui {
	return e.(uint)
}

func AsUi8(e Any) Ui8 {
	switch {
	case IsI(e):
		return Ui8(AsI(e))
	case IsI8(e):
		return Ui8(AsI8(e))
	case IsI16(e):
		i16 := AsI16(e)
		if i16 <= MaxInt8 {
			return Ui8(i16)
		}
		return MaxInt8
	case IsI32(e):
		i32 := AsI32(e)
		if i32 <= MaxInt8 {
			return Ui8(i32)
		}
		return MaxInt8
	case IsI64(e):
		i64 := AsI64(e)
		if i64 <= MaxInt8 {
			return Ui8(i64)
		}
		return MaxInt8
	default:
		return MaxInt8
	}
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

func AsI(e Any) I {
	switch {
	case IsI8(e):
		return I(AsI8(e))
	case IsI16(e):
		return I(AsI16(e))
	case IsI32(e):
		return I(AsI32(e))
	case IsI64(e):
		return I(AsI64(e))
	case IsI(e):
		return e.(I)
	default:
		return 0
	}
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
