package hftypes

import "reflect"

func AnyToSliceAny(es Any) Anys {
	res := Anys{}
	rv := reflect.ValueOf(es)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			res = append(res, rv.Index(i).Interface())
		}
	}
	return res
}
