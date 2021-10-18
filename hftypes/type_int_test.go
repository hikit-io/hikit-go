package hftypes

import "testing"

func TestAsI64(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    I64
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			want: MaxInt64,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			want: MaxInt32,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			want: MaxInt64,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			want: MaxUint8,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			want: MaxUint16,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			want: MaxUint32,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
			recover: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.recover {
				defer func() {
					if err := recover(); err != nil {

					}
				}()
				if got := AsI64(tt.args.e); got != tt.want {
					t.Errorf("AsI8() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsI64(tt.args.e); got != tt.want {
				t.Errorf("AsI64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsI32(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    I32
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int(1000),
			},
			want: 1000,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			want: MaxInt32,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int64(2000),
			},
			want: 2000,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			want: MaxUint8,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			want: MaxUint16,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint32(3000),
			},
			want: 3000,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
			recover: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.recover {
				defer func() {
					if err := recover(); err != nil {

					}
				}()
				if got := AsI32(tt.args.e); got != tt.want {
					t.Errorf("AsI8() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsI32(tt.args.e); got != tt.want {
				t.Errorf("AsI32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsI16(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    I16
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int(1000),
			},
			want: 1000,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int64(2000),
			},
			want: 2000,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			want: MaxUint8,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint32(3000),
			},
			want: 3000,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
			recover: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.recover {
				defer func() {
					if err := recover(); err != nil {

					}
				}()
				if got := AsI16(tt.args.e); got != tt.want {
					t.Errorf("AsI8() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsI16(tt.args.e); got != tt.want {
				t.Errorf("AsI16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsI8(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    I8
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int(100),
			},
			want: 100,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int16(123),
			},
			want: 123,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: int64(124),
			},
			want: 124,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(125),
			},
			want: 125,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint32(126),
			},
			want: 126,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
			recover: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.recover {
				defer func() {
					if err := recover(); err != nil {

					}
				}()
				if got := AsI8(tt.args.e); got != tt.want {
					t.Errorf("AsI8() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsI8(tt.args.e); got != tt.want {
				t.Errorf("AsI8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsUi64(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    Ui64
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			want: MaxInt64,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			want: MaxInt32,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			want: MaxInt64,
		},
		{
			name: "",
			args: args{
				e: int64(MinInt64),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			want: MaxUint64,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			want: MaxUint8,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			want: MaxUint16,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			want: MaxUint32,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			want: MaxUint64,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
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
				if got := AsUi64(tt.args.e); got != tt.want {
					t.Errorf("AsUi64() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsUi64(tt.args.e); got != tt.want {
				t.Errorf("AsUi64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsUi32(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    Ui32
		recover B
	}{
		{
			name: "",
			args: args{
				e: int(MaxInt64),
			},
			recover: true,
			//want: MaxInt64,
		},
		{
			name: "",
			args: args{
				e: int8(MaxInt8),
			},
			want: MaxInt8,
		},
		{
			name: "",
			args: args{
				e: int16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: int32(MaxInt32),
			},
			want: MaxInt32,
		},
		{
			name: "",
			args: args{
				e: int64(MaxInt64),
			},
			//want: MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: uint8(MaxUint8),
			},
			want: MaxUint8,
		},
		{
			name: "",
			args: args{
				e: uint16(MaxUint16),
			},
			want: MaxUint16,
		},
		{
			name: "",
			args: args{
				e: uint32(MaxUint32),
			},
			want: MaxUint32,
		},
		{
			name: "",
			args: args{
				e: uint64(MaxUint64),
			},
			//want: MaxUint64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "string",
			},
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
				if got := AsUi32(tt.args.e); got != tt.want {
					t.Errorf("AsUi32() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsUi32(tt.args.e); got != tt.want {
				t.Errorf("AsUi32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsUi8(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    Ui8
		recover B
	}{
		{
			name: "",
			args: args{
				e: "asdsad",
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: Ui8(10),
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				e: I(10),
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				e: I(513),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: I16(123),
			},
			want: 123,
		},
		{
			name: "",
			args: args{
				e: I16(3123),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: I32(123),
			},
			want: 123,
		},
		{
			name: "",
			args: args{
				e: I32(123213),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: I64(123),
			},
			want: 123,
		},
		{
			name: "",
			args: args{
				e: I64(123123),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: Ui16(133),
			},
			want: 133,
		},
		{
			name: "",
			args: args{
				e: Ui16(12313),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: Ui64(12313),
			},
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
				if got := AsUi8(tt.args.e); got != tt.want {
					t.Errorf("AsUi8() = %v, want %v", got, tt.want)
				}
			}
			if got := AsUi8(tt.args.e); got != tt.want {
				t.Errorf("AsUi8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsI(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    I
		recover B
	}{
		{
			name: "",
			args: args{
				e: I(12),
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				e: I32(12),
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				e: I64(MaxInt64),
			},
			want:    MaxInt64,
			recover: true,
		},
		{
			name: "",
			args: args{
				e: Ui(12),
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				e: Ui(12),
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				e: Ui32(MaxUint32),
			},
			want: MaxUint32,
		},
		{
			name: "",
			args: args{
				e: Ui64(MaxUint64),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "asdsad",
			},
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
				if got := AsI(tt.args.e); got != tt.want {
					t.Errorf("AsI() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsI(tt.args.e); got != tt.want {
				t.Errorf("AsI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsUi16(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name    string
		args    args
		want    Ui16
		recover B
	}{
		{
			name: "",
			args: args{
				e: Ui16(120),
			},
			want: 120,
		},
		{
			name: "",
			args: args{
				e: Ui16(MaxUint16),
			},
			want: MaxUint16,
		},
		{
			name: "",
			args: args{
				e: Ui32(3123),
			},
			want: 3123,
		},
		{
			name: "",
			args: args{
				e: Ui32(MaxUint32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: Ui64(12132),
			},
			want: 12132,
		},
		{
			name: "",
			args: args{
				e: I(MinInt8),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: I16(MaxInt16),
			},
			want: MaxInt16,
		},
		{
			name: "",
			args: args{
				e: I32(MaxInt32),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: I64(MaxInt64),
			},
			recover: true,
		},
		{
			name: "",
			args: args{
				e: "sadsad",
			},
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
				if got := AsUi16(tt.args.e); got != tt.want {
					t.Errorf("AsUi16() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := AsUi16(tt.args.e); got != tt.want {
				t.Errorf("AsUi16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsI64sOrPtr(t *testing.T) {
	type args struct {
		elem Any
	}
	tests := []struct {
		name string
		args args
		want B
	}{
		{
			name: "",
			args: args{
				elem: []int{},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				elem: []int64{},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				elem: []*int64{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsI64sOrPtr(tt.args.elem); got != tt.want {
				t.Errorf("IsI64sOrPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
