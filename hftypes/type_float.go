package hftypes

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
	return e.(F32)
}

func AsF64(e Any) F64 {
	return e.(F64)
}
