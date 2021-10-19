package hftypes

import (
	"fmt"
	"testing"
)

func TestStructToMapStrAny(t *testing.T) {
	type AddrItem struct {
		Sex string
	}
	type Addr struct {
		Street string
		Post   *string
		Item   *AddrItem
	}
	type User struct {
		Name string
		Age  int
		Addr Addr
		Ptr  *int
	}
	m := map[Str]Any{}
	u := User{
		Name: "hfunc",
		Age:  0,
		Addr: Addr{},
		Ptr:  nil,
	}
	StructToMapStrAny(u, m)
	fmt.Println(m)
	u.Ptr = IPtr(1)
	u.Addr.Post = StrPtr("das")
	u.Addr.Item = &AddrItem{Sex: "dasd"}
	StructToMapStrAny(u, &m)
	fmt.Println(m)
}

func TestStructToMapStrAny1(t *testing.T) {
	type AddrItem struct {
		Sex string
	}
	type Addr struct {
		Street string
		Post   *string
		Item   *AddrItem
	}
	type User struct {
		Name string
		Age  int
		Addr Addr
		Ptr  *int
	}
	type args struct {
		s MustStruct
		m MustMapStrAnyOrPtr
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: User{},
				m: MapStrAny{},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: User{},
				m: &MapStrAny{},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: &User{},
				m: MapStrAny{},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: &User{},
				m: &MapStrAny{},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: "",
				m: nil,
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				s: nil,
				m: MapStrAny{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StructToMapStrAny(tt.args.s, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("StructToMapStrAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
