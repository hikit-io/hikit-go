package hfmongo

import (
	"fmt"
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

func Test_mergeField(t *testing.T) {
	op := Builder{}
	op.Field("name").Equal("NieAowei")
	op.Field("name").Equal("NieAowei1")
	op.Field("name").LessThan(10)
	op.Field("age").Equal(18)
	op.Field("name").Child("first").Equal("Nie")
	op.Field("name").Child("first").Child("one").Equal("Nie")

	op.Field("name").Set("nieaowei")
	op.Field("age").Set(100)

	op.Field("name").Child("first").Set("nekilc")
	op.Field("name").Child("first").Child("one").Set("Nie")
	orBd := NewBuilder()
	orBd.Field("age").GreatThan(10).LessThan(20)
	op.Or(orBd)
	op.OrFc(func(b *Builder) {
		b.Field("address").Equal("shanghai")
	})
	//res := mergeFindField("", op.fields)
	fmt.Println(op.Filter().Map())
	//fmt.Println(mergeUpField("", op.fields).Map())
	fmt.Println(op.Update().Map())
}

func TestModel(t *testing.T) {
	type M struct {
		Model
	}
	var m interface{} = &M{}
	fmt.Println(m.(*Model))
}
