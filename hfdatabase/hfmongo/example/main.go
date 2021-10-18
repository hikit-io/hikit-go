package main

import (
	"context"
	"github.com/hfunc/hfunc-go/hfdatabase/hfmongo"
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

func init() {
	ctx := context.Background()
	o := options.Client()
	o.ApplyURI(url)
	cli, _ := mongo.NewClient(o)
	cli.Connect(ctx)
	db := cli.Database("test")
	col = db.Collection("test")
}

func main() {

	builder := &hfmongo.Builder{}
	builder.Field("name").Regex("nieaowei")
	builder.Field("age").LessThan(30)
	builder.Field("addr").Projection(false)
	builder.Limit(2)
	builder.Skip(1)
	op := builder.Options()
	cur, err := col.Find(ctx, builder.Filter(), op)
	if err != nil {
		panic(err)
	}
	users := []User{}
	err = cur.All(ctx, &users)
	if err != nil {
		panic(err)
	}
}
