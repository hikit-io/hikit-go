package hfgin

import "testing"

func TestMatchUrl(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				"Hello_idWei",
			},
			want: "/hello_id/wei",
		},
		{
			name: "",
			args: args{
				"Hello_IdWei",
			},
			want: "/hello/:id/wei",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchUrl(tt.args.name); got != tt.want {
				t.Errorf("MatchUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
