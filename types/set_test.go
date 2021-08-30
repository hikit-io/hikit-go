package types

import "testing"

func TestSet(t *testing.T) {
	set := NewSet(String)
	set.Add("das")
	set.Del("das")
}
