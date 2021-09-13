package hftypes

func IsBool(elem Any) B {
	switch elem.(type) {
	case B, *B:
		return true
	}
	return false
}

var IsBools = IsBs

func IsBs(elems Any) B {
	switch elems.(type) {
	case []B, []*B:
		return true
	}
	return false
}

func IsB(elem Any) B {
	switch elem.(type) {
	case B, *B:
		return true
	}
	return false
}

func ToBool(e Any) B {
	switch {
	case IsInt(e):
		AsI64(e)
	}
	return e.(B)
}

func AsBool(e Any) B {
	return e.(bool)
}
