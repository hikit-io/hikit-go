package hftypes

import (
	"fmt"
	"strconv"
)

func IsString(elem Any) B {
	switch elem.(type) {
	case string, *string:
		return true
	}
	return false
}

func IsStr(elem Any) B {
	_, ok := elem.(string)
	return ok
}

func IsStrp(elem Any) B {
	_, ok := elem.(*string)
	return ok
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
