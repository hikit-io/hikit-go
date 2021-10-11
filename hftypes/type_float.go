package hftypes

const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

func RangeF32(f F64) B {
	return f <= MaxFloat32 && f >= SmallestNonzeroFloat32
}

func RangeF64(f F64) B {
	return f <= MaxFloat64 && f >= SmallestNonzeroFloat64
}

func IRangeF64(f I64) B {
	return F64(f) < MaxFloat64 && F64(f) > SmallestNonzeroFloat64
}

func URangeF64(f Ui64) B {
	return F64(f) < MaxFloat64 && F64(f) > SmallestNonzeroFloat64
}

func F32Ptr(f F32) *F32 {
	return &f
}

func F64Ptr(f F64) *F64 {
	return &f
}

func IsFloat(elem Any) B {
	switch elem.(type) {
	case F32, F64, *F32, *F64:
		return true
	}
	return false
}

func IsF32(elem Any) B {
	_, ok := elem.(F32)
	return ok
}

func IsF64(elem Any) B {
	_, ok := elem.(F64)
	return ok
}

func AsF32(e Any) F32 {
	switch {
	case IsF32(e):
		return e.(F32)
	case IsF64(e):
		f64 := AsF64(e)
		if RangeF32(f64) {
			return F32(f64)
		}
		panic(f64)
	default:
		panic("not support")
	}
}

func AsF64(e Any) F64 {
	switch {
	case IsF32(e):
		return F64(AsF32(e))
	case IsF64(e):
		return e.(F64)
	case IsInt(e):
		i64 := AsI64(e)
		return F64(i64)
	default:
		panic("not support")
	}
}
