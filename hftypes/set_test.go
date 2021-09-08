package hftypes

import "testing"

func TestSet(t *testing.T) {
	set := Set{}
	set.Init()
	set.Add("das")
	set.Add("das23")
	set.Del("das")
	t.Log(set.Strings())
}
