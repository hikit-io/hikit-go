package hfmongo

import (
	"testing"
)

func TestOp(t *testing.T) {
	//col := (&mongo.Client{}).Database("d").Collection("s")
	op := Builder{}
	op.Field("name").In([]string{"123", "ewq"})
	t.Log(op.Filter())
	op.Field("name").In([]string{"123", "ewq", "dsad"}).Equal("ad")
	op.Field("name1").In([]string{"123", "ewq", "dsad"})
	t.Log(op.Filter())
	op.Field("name").Set([]string{"123", "ewq"}).UnSet()
	t.Log(op.Filter())
}
