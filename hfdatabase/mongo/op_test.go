package mongo

import (
	"testing"
)

func TestOp(t *testing.T) {
	op := Op{}
	op.Field("name").In([]string{"123", "ewq"})
	t.Log(op.Find())
	op.Field("name").In([]string{"123", "ewq", "dsad"}).Eq("ad")
	op.Field("name1").In([]string{"123", "ewq", "dsad"})
	t.Log(op.Find())
	op.Field("name").Set([]string{"123", "ewq"}).UnSet()
	t.Log(op.Update())
}
