package hftypes

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	set := Set{}
	set.Init()
	set.Add("das")
	set.Add("das23")
	set.Add("das24")
	t.Run("", func(t *testing.T) {
		if set.Size() != 3 {
			t.Errorf("size != 3")
		}
	})
	set.Add("das24")
	t.Run("repeat", func(t *testing.T) {
		if set.Size() != 3 {
			t.Errorf("size != 3")
		}
	})
	set.Del("das24")
	t.Run("del", func(t *testing.T) {
		if set.Size() != 2 {
			t.Errorf("size != 3")
		}
	})
	t.Run("strings slice", func(t *testing.T) {
		if !reflect.DeepEqual(set.Strings(), []string{
			"das", "das23",
		}) {
			t.Errorf("string slice ! %+v", set.Strings())
		}
	})
	t.Run("exist true", func(t *testing.T) {
		if !set.Exist("das") {
			t.Errorf("")
		}
	})

	t.Run("exist false", func(t *testing.T) {
		if set.Exist("das123") {
			t.Errorf("")
		}
	})
	t.Run("exist err", func(t *testing.T) {
		if set.Exist(123) {
			t.Errorf("")
		}
	})
	setInt := NewSet(IsInt)
	t.Run("not suppory type", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {

			}
		}()
		if setInt.Add("dsad") {
			t.Errorf("")
			return
		}
	})
}

func TestSet_Strings(t *testing.T) {
	tests := []struct {
		name   string
		fields *Set
		want   []Str
		initFc func(set *Set)
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: NewSet(IsStr),
			want: []string{
				"str", "str1", "str2",
			},
			initFc: func(set *Set) {
				set.Add("str")
				set.Add("str1")
				set.Add("str2")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initFc(tt.fields)
			got := tt.fields.Strings()
			count := 0
			for _, g := range got {
				for _, str := range tt.want {
					if g == str {
						count++
						break
					}
				}
			}
			if count != tt.fields.Size() {
				t.Errorf("Strings() = %v, want %v", count, tt.fields.Size())
			}
		})
	}
}

func TestSet_Exist(t *testing.T) {
	type fields struct {
		data   map[Any]Void
		TypeFc func(elem Any) B
		init   B
	}
	d := fields{
		data:   map[Any]Void{},
		TypeFc: IsStr,
		init:   false,
	}
	type args struct {
		elem Any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   B
	}{
		{
			name:   "",
			fields: d,
			args: args{
				elem: "q",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				data:   tt.fields.data,
				TypeFc: tt.fields.TypeFc,
				init:   tt.fields.init,
			}
			if got := s.Exist(tt.args.elem); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
