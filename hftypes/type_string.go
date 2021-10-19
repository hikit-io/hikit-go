package hftypes

import (
	"fmt"
	"strconv"
)

func IsStr(elem Any) B {
	_, ok := elem.(string)
	return ok
}

func IsStrPtr(elem Any) B {
	_, ok := elem.(*string)
	return ok
}

func IsStrOrPtr(elem Any) B {
	if IsStr(elem) {
		return true
	}
	return IsStrPtr(elem)
}

func IsStrs(elem Any) B {
	_, ok := elem.([]string)
	return ok
}

func IsStrsPtr(elem Any) B {
	_, ok := elem.([]*string)
	return ok
}

func IsStrsOrPtr(elem Any) B {
	if IsStrs(elem) {
		return true
	}
	return IsStrsPtr(elem)
}

func ToStr(e Any) string {
	switch {
	case IsBool(e):
		return strconv.FormatBool(AsBool(e))
	case IsStr(e):
		return AsStr(e)
	case IsInt(e):
		return strconv.FormatInt(AsI64(e), 10)
	case IsUint(e):
		return strconv.FormatUint(AsUi64(e), 10)
	case IsFloat(e):
		return strconv.FormatFloat(AsF64(e), 10, 10, 64)
	}
	return fmt.Sprintf("%+v", e)
}

func AsStr(e Any) string {
	return e.(string)
}

func StrPtr(s Str) *Str {
	return &s
}
