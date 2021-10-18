package hftypes

import (
	"fmt"
	"testing"
)

func TestStructToMapStrAny(t *testing.T) {
	type Addr struct {
		Street string
	}
	type User struct {
		Name string
		Age  int
		Addr Addr
	}
	m := map[Str]Any{}
	StructToMapStrAny(User{}, m)
	fmt.Println(m)
	StructToMapStrAny(User{}, &m)
	fmt.Println(m)
}
