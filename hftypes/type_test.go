package hftypes

import "testing"

func TestAsInt32(t *testing.T) {

}

func TestIsArray(t *testing.T) {
	type args struct {
		elem Any
	}
	tests := []struct {
		name string
		args args
		want B
	}{
		{
			name: "common",
			args: args{
				elem: 9,
			},
			want: false,
		},
		{
			name: "array",
			args: args{
				elem: []string{"dsa"},
			},
			want: true,
		},
		{
			name: "array ptr",
			args: args{
				elem: &([]string{"dsa"}),
			},
			want: true,
		},
		{
			name: "array struct",
			args: args{
				elem: &([]struct{ Name string }{{"dsa"}}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlice(tt.args.elem); got != tt.want {
				t.Errorf("IsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsBool(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    B
		recover B
	}{
		// TODO: Add test cases.
		{
			name: "bool",
			args: args{
				e: true,
			},
			want: true,
		},
		{
			name: "bool",
			args: args{
				e: "true",
			},
			want:    false,
			recover: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if tt.recover {
				defer func() {
					if r := recover(); r != nil {

					}
				}()
				if got := AsBool(tt.args.e); got != tt.want {
					t.Errorf("AsBool() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsBool(tt.args.e); got != tt.want {
				t.Errorf("AsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestAsI64(t *testing.T) {
//	type args struct {
//		e Any
//	}
//	tests := []struct {
//		name string
//		args args
//		want I64
//	}{
//		{
//			name: "int",
//			args: args{
//				e: int(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int8",
//			args: args{
//				e: int8(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int16",
//			args: args{
//				e: int16(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int32",
//			args: args{
//				e: int32(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 short",
//			args: args{
//				e: int64(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 long",
//			args: args{
//				e: int64(8446462598732840960),
//			},
//			want: 8446462598732840960,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := AsI64(tt.args.e); got != tt.want {
//				t.Errorf("AsI64() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestAsI(t *testing.T) {
//	type args struct {
//		e Any
//	}
//	tests := []struct {
//		name string
//		args args
//		want I
//	}{
//		{
//			name: "int",
//			args: args{
//				e: int(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int8",
//			args: args{
//				e: int8(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int16",
//			args: args{
//				e: int16(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int32",
//			args: args{
//				e: int32(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 short",
//			args: args{
//				e: int64(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 long",
//			args: args{
//				e: int64(8446462598732840960),
//			},
//			want: 8446462598732840960,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := AsI(tt.args.e); got != tt.want {
//				t.Errorf("AsI() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestAsI32(t *testing.T) {
//	type args struct {
//		e Any
//	}
//	tests := []struct {
//		name string
//		args args
//		want I32
//	}{
//		{
//			name: "int",
//			args: args{
//				e: int(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int8",
//			args: args{
//				e: int8(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int16",
//			args: args{
//				e: int16(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int32",
//			args: args{
//				e: int32(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 short",
//			args: args{
//				e: int64(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 long",
//			args: args{
//				e: int64(8446462598732840960),
//			},
//			want: MaxInt32,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := AsI32(tt.args.e); got != tt.want {
//				t.Errorf("AsI32() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestAsI8(t *testing.T) {
//	type args struct {
//		e Any
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    I8
//		recover B
//	}{
//		{
//			name: "int",
//			args: args{
//				e: int(10),
//			},
//			want: 10,
//		},
//		{
//			name: "uint",
//			args: args{
//				e: uint(10),
//			},
//			want: 10,
//		},
//		{
//			name: "uint>max8",
//			args: args{
//				e: uint(MaxInt8 + 1),
//			},
//			recover: true,
//		},
//		{
//			name: "int>max8",
//			args: args{
//				e: int(MaxInt8 + 1),
//			},
//			want:    10,
//			recover: true,
//		},
//		{
//			name: "int8",
//			args: args{
//				e: int8(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int16",
//			args: args{
//				e: int16(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int16>max8",
//			args: args{
//				e: int16(MaxInt8 + 1),
//			},
//			recover: true,
//			want:    10,
//		},
//		{
//			name: "int32",
//			args: args{
//				e: int32(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int32>max8",
//			args: args{
//				e: int32(MaxInt8 + 1),
//			},
//			recover: true,
//		},
//		{
//			name: "int64 short",
//			args: args{
//				e: int64(10),
//			},
//			want: 10,
//		},
//		{
//			name: "int64 long",
//			args: args{
//				e: int64(8446462598732840960),
//			},
//			want:    0,
//			recover: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if tt.recover {
//				defer func() {
//					if err := recover(); err != nil {
//
//					}
//				}()
//				if got := AsI8(tt.args.e); got != tt.want {
//					t.Errorf("AsI8() = %v, want %v", got, tt.want)
//				}
//				return
//			}
//			if got := AsI8(tt.args.e); got != tt.want {
//				t.Errorf("AsI8() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
