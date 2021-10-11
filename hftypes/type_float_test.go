package hftypes

import "testing"

func TestAsF64(t *testing.T) {
	type args struct {
		e Any
	}
	tests := []struct {
		name string
		args args
		want F64
	}{
		{
			name: "int",
			args: args{
				e: int(10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsF64(tt.args.e); got != tt.want {
				t.Errorf("AsF64() = %v, want %v", got, tt.want)
			}
		})
	}
}
