package hktypes

func IPtr(i I) *I {
	return &i
}

func I8Ptr(i I8) *I8 {
	return &i
}

func I16Ptr(i I16) *I16 {
	return &i
}

func I32Ptr(i I32) *I32 {
	return &i
}

func I64Ptr(i I64) *I64 {
	return &i
}

func UiPtr(i Ui) *Ui {
	return &i
}

func Ui8Ptr(i Ui8) *Ui8 {
	return &i
}

func Ui16Ptr(i Ui16) *Ui16 {
	return &i
}

func Ui32Ptr(i Ui32) *Ui32 {
	return &i
}

func Ui64Ptr(i Ui64) *Ui64 {
	return &i
}

func IsUi(elem Any) B {
	_, ok := elem.(uint)
	return ok
}

func IsUiPtr(elem Any) B {
	_, ok := elem.(*uint)
	return ok
}

func IsUi8(elem Any) B {
	_, ok := elem.(uint8)
	return ok
}

func IsUi8Ptr(elem Any) B {
	_, ok := elem.(*Ui8)
	return ok
}

func IsUi16(elem Any) B {
	_, ok := elem.(uint16)
	return ok
}

func IsUi16Ptr(elem Any) B {
	_, ok := elem.(*uint16)
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

func IsUi64Ptr(elem Any) B {
	_, ok := elem.(*uint64)
	return ok
}

func IsI(elem Any) B {
	_, ok := elem.(int)
	return ok
}

func IsIPtr(elem Any) B {
	_, ok := elem.(*int)
	return ok
}

func IsI8(elem Any) B {
	_, ok := elem.(int8)
	return ok
}

func IsI8Ptr(elem Any) B {
	_, ok := elem.(*int8)
	return ok
}

func IsI16(elem Any) B {
	_, ok := elem.(int16)
	return ok
}

func IsI16Ptr(elem Any) B {
	_, ok := elem.(*int16)
	return ok
}

func IsI32(elem Any) B {
	_, ok := elem.(int32)
	return ok
}

func IsI32Ptr(elem Any) B {
	_, ok := elem.(*int32)
	return ok
}

func IsI64(elem Any) B {
	_, ok := elem.(int64)
	return ok
}

func IsI64Ptr(elem Any) B {
	_, ok := elem.(*int64)
	return ok
}

func IsInt(elem Any) B {
	switch elem.(type) {
	case I, I8, I16, I32, I64:
		return true
	}
	return false
}

func IsIntPtr(e Any) B {
	switch e.(type) {
	case I, I8, I16, I32, I64:
		return true
	}
	return false
}

func IsIntOrPtr(e Any) B {
	switch e.(type) {
	case I, I8, I16, I32, I64, *I, *I8, *I16, *I32, *I64:
		return true
	}
	return false
}

func IsUint(elem Any) B {
	switch elem.(type) {
	case Ui, Ui8, Ui16, Ui32, Ui64:
		return true
	}
	return false
}

func IsUintPtr(elem Any) B {
	switch elem.(type) {
	case *Ui, *Ui8, *Ui16, *Ui32, *Ui64:
		return true
	}
	return false
}

func IsUintOrPtr(e Any) B {
	switch e.(type) {
	case Ui, Ui8, Ui16, Ui32, Ui64, *Ui, *Ui8, *Ui16, *Ui32, *Ui64:
		return true
	}
	return false
}

func IsIs(elem Any) B {
	_, ok := elem.([]int)
	return ok
}

func IsIsPtr(elem Any) B {
	_, ok := elem.([]*int)
	return ok
}

func IsIsOrPtr(elem Any) B {
	if IsIs(elem) {
		return true
	}
	return IsIsPtr(elem)
}

func IsI8s(elem Any) B {
	_, ok := elem.([]I8)
	return ok
}

func IsI8sPtr(elem Any) B {
	_, ok := elem.([]*I8)
	return ok
}

func IsI8sOrPtr(elem Any) B {
	if IsI8s(elem) {
		return true
	}
	return IsI8sPtr(elem)
}

func IsI16s(elem Any) B {
	_, ok := elem.([]I16)
	return ok
}

func IsI16sPtr(elem Any) B {
	_, ok := elem.([]*I16)
	return ok
}

func IsI16sOrPtr(elem Any) B {
	if IsI16s(elem) {
		return true
	}
	return IsI16sPtr(elem)
}

func IsI32s(elem Any) B {
	_, ok := elem.([]I32)
	return ok
}

func IsI32sPtr(elem Any) B {
	_, ok := elem.([]*I32)
	return ok
}

func IsI32sOrPtr(elem Any) B {
	if IsI32s(elem) {
		return true
	}
	return IsI32sPtr(elem)
}

func IsI64s(elem Any) B {
	_, ok := elem.([]I64)
	return ok
}

func IsI64sPtr(elem Any) B {
	_, ok := elem.([]*I64)
	return ok
}

func IsI64sOrPtr(elem Any) B {
	if IsI64s(elem) {
		return true
	}
	return IsI64sPtr(elem)
}

func IsUis(elem Any) B {
	_, ok := elem.([]Ui)
	return ok
}

func IsUisPtr(elem Any) B {
	_, ok := elem.([]*Ui)
	return ok
}

func IsUisOrPtr(elem Any) B {
	if IsUis(elem) {
		return true
	}
	return IsUisPtr(elem)
}

func IsUi8s(elem Any) B {
	_, ok := elem.([]Ui8)
	return ok
}

func IsUi8sPtr(elem Any) B {
	_, ok := elem.([]*Ui8)
	return ok
}

func IsUi8sOrPtr(elem Any) B {
	if IsUi8s(elem) {
		return true
	}
	return IsUi8sPtr(elem)
}

func IsUi16s(elem Any) B {
	_, ok := elem.([]Ui16)
	return ok
}

func IsUi16sPtr(elem Any) B {
	_, ok := elem.([]*Ui16)
	return ok
}

func IsUi16sOrPtr(elem Any) B {
	if IsUis(elem) {
		return true
	}
	return IsUisPtr(elem)
}

func IsUi32s(elem Any) B {
	_, ok := elem.([]Ui32)
	return ok
}

func IsUi32sPtr(elem Any) B {
	_, ok := elem.([]*Ui)
	return ok
}

func IsUi32sOrPtr(elem Any) B {
	if IsUi32s(elem) {
		return true
	}
	return IsUi32sPtr(elem)
}

func IsUi64s(elem Any) B {
	_, ok := elem.([]Ui64)
	return ok
}

func IsUi64sPtr(elem Any) B {
	_, ok := elem.([]*Ui64)
	return ok
}

func IsUi64sOrPtr(elem Any) B {
	if IsUis(elem) {
		return true
	}
	return IsUisPtr(elem)
}

func IsF32s(elem Any) B {
	_, ok := elem.([]F32)
	return ok
}

func IsF32sPtr(elem Any) B {
	_, ok := elem.([]*F32)
	return ok
}

func IsF32sOrPtr(elem Any) B {
	if IsF32(elem) {
		return true
	}
	return IsF32sPtr(elem)
}

func IsF64s(elem Any) B {
	_, ok := elem.([]F64)
	return ok
}

func IsF64sPtr(elem Any) B {
	_, ok := elem.([]*F32)
	return ok
}

func IsF64sOrPtr(elem Any) B {
	if IsF64s(elem) {
		return true
	}
	return IsF64sPtr(elem)
}

func RangeI8(i64 I64) bool {
	return i64 >= MinInt8 && i64 <= MaxInt8
}

func URangeI8(ui64 Ui64) bool {
	return ui64 <= MaxInt8
}

func RangeUi8(i64 I64) bool {
	return i64 >= 0 && i64 <= MaxUint8
}

func URangeUi8(ui64 Ui64) bool {
	return ui64 <= MaxUint8
}

func RangeI16(i64 I64) bool {
	return i64 >= MinInt16 && i64 <= MaxInt16
}

func URangeI16(ui64 Ui64) bool {
	return ui64 <= MaxInt16
}

func RangeUi16(i64 I64) bool {
	return i64 >= 0 && i64 <= MaxUint16
}

func URangeUi16(ui64 Ui64) bool {
	return ui64 <= MaxUint16
}

func RangeI32(i64 I64) bool {
	return i64 >= MinInt32 && i64 <= MaxInt32
}

func URangeI32(ui64 Ui64) bool {
	return ui64 <= MaxInt32
}

func RangeUi32(i64 I64) bool {
	return i64 >= 0 && i64 <= MaxUint32
}

func URangeUi32(ui64 Ui64) bool {
	return ui64 <= MaxUint32
}

func RangeI64(i64 I64) bool {
	return i64 >= MinInt64 && i64 <= MaxInt64
}

func URangeI64(ui64 Ui64) bool {
	return ui64 <= MaxInt64
}

func RangeUi64(i64 I64) bool {
	return i64 >= 0
}

func URangeUi64(ui64 Ui64) bool {
	return ui64 <= MaxUint64
}

func AsI8(e Any) I8 {
	switch {
	case IsI8(e):
		return e.(I8)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeI8(i64) {
			return I8(i64)
		}
		panic(i64)
	case IsUint(e):
		i := AsUi64(e)
		if URangeI8(i) {
			return I8(i)
		}
		panic(i)
	default:
		panic("not support")
	}
}

func AsI16(e Any) I16 {
	switch {
	case IsI16(e):
		return e.(I16)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeI16(i64) {
			return I16(i64)
		}
		panic(i64)
	case IsUint(e):
		i := AsUi64(e)
		if URangeI16(i) {
			return I16(i)
		}
		panic(i)
	default:
		panic("not support")
	}
}

func AsI32(e Any) I32 {
	switch {
	case IsI32(e):
		return e.(I32)
	case IsUint(e):
		i := AsUi64(e)
		if URangeI32(i) {
			return I32(i)
		}
		panic(i)
	case IsInt(e):
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
	case IsUint(e):
		ui64 := AsUi64(e)
		if URangeI64(ui64) {
			return I64(ui64)
		}
		panic(ui64)
	default:
		panic("not support type")
	}
}

func AsUi(e Any) Ui {
	return e.(uint)
}

func AsUi8(e Any) Ui8 {
	switch {
	case IsUi8(e):
		return e.(Ui8)
	case IsUint(e):
		ui64 := AsUi64(e)
		if URangeUi8(ui64) {
			return Ui8(ui64)
		}
		panic(ui64)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeUi8(i64) {
			return Ui8(i64)
		}
		panic(i64)
	default:
		panic("not support type")
	}
}

func AsUi16(e Any) Ui16 {
	switch {
	case IsUi16(e):
		return e.(Ui16)
	case IsUint(e):
		ui64 := AsUi64(e)
		if URangeUi16(ui64) {
			return Ui16(ui64)
		}
		panic(ui64)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeUi16(i64) {
			return Ui16(i64)
		}
		panic(i64)
	default:
		panic("not support")
	}
}

func AsUi32(e Any) Ui32 {
	switch {
	case IsUi32(e):
		return e.(Ui32)
	case IsUint(e):
		ui64 := AsUi64(e)
		if URangeUi32(ui64) {
			return Ui32(ui64)
		}
		panic(ui64)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeUi32(i64) {
			return Ui32(i64)
		}
		panic(i64)
	default:
		panic("not support")
	}
}

func AsUi64(e Any) Ui64 {
	switch {
	case IsUi(e):
		return Ui64(AsUi(e))
	case IsUi8(e):
		return Ui64(AsUi8(e))
	case IsUi16(e):
		return Ui64(AsUi16(e))
	case IsUi32(e):
		return Ui64(AsUi32(e))
	case IsUi64(e):
		return e.(Ui64)
	case IsInt(e):
		i := AsI64(e)
		if RangeUi64(i) {
			return Ui64(i)
		}
		panic(i)
	}
	panic(e)
}

func RangeI(i64 I64) B {
	if IntSize == 64 {
		return i64 <= MaxInt64
	} else {
		return i64 <= MaxInt32
	}
}

func URangeI(i64 Ui64) B {
	if IntSize == 64 {
		return i64 <= MaxInt64
	} else {
		return i64 <= MaxInt32
	}
}

func AsI(e Any) I {
	switch {
	case IsI(e):
		return e.(I)
	case IsInt(e):
		i64 := AsI64(e)
		if RangeI(i64) {
			return I(i64)
		}
		panic(i64)
	case IsUint(e):
		ui64 := AsUi64(e)
		if URangeI(ui64) {
			return I(ui64)
		}
		panic(ui64)
	default:
		panic(e)
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
