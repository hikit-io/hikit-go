package hftypes

import "testing"

func TestAsInt32(t *testing.T) {
	i32 := int32(12)
	i64 := (I64)(i32)
	AsI64(i32)
	AsI16(i32)
	AsI(i64)
}
