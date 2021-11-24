package main

import (
	"context"
	"fmt"
	"go.hikit.io/database/hkmg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  uint64             `bson:"age"`
	Addr string             `bson:"addr"`
}

var col *mongo.Collection
var ctx = context.Background()

var db *hkmg.Database

func init() {
	ctx := context.Background()
	o := options.Client()
	o.ApplyURI(url)
	cli, _ := mongo.NewClient(o)
	err := cli.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db = hkmg.NewDB(cli, "test")
	//col = db.Collection("test")

}

type Test struct {
	Name string `json:"name" bson:"name"`
	Age  uint8  `json:"age" bson:"age"`
	Addr string `json:"addr" bson:"addr"`
}
type TestAge struct {
	Age int32 `bson:"age"`
}

func (t Test) DocName() string {
	return "test"
}

func main() {
	e := db.Col(Test{}).HInsertOne(ctx, &Test{
		Age: 111,
	})
	if e.NotNil() {
		panic(e.Err())
	}

	builder := hkmg.NewBuilder()
	//builder.Field("name").Regex("nieaowei")
	//builder.Field("age").LessThan(30)
	builder.OrFc(func(br *hkmg.Builder) {
		br.Field("name").Equal("nieaowei")
	}).OrFc(func(br *hkmg.Builder) {
		br.Field("age").Equal(11)
	})
	ts := []TestAge{}

	err := db.Col(User{}).HFind(ctx, builder, &ts)
	if err.Err() != nil {
		panic(err)
	}
	fmt.Println(ts)

	querySturct := &Test{
		Age: 11,
	}

	findRes := db.Col(User{}).HFind(ctx, querySturct, &ts)
	if findRes.Err() != nil {
		panic(err)
	}
	fmt.Println(ts)

	findOneRes := db.Col(User{}).HFindOne(ctx, builder, querySturct)
	if findOneRes.Err() != nil {
		panic(err)
	}
	fmt.Println(querySturct)

	r := db.Col(User{}).HUpdateOne(ctx, builder, querySturct)
	if r.Err() != nil {
		panic(r.Error())
	}

	if r.ExceptNoDocuments() != nil {

	}
	fmt.Println(querySturct)
}
