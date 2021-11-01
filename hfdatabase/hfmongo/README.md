# hfmongo

## Features

- Chain API
- Custom log

## Getting start

`go get github.com/hfunc/hfunc-go/hfdatabase/hfmongo`

### Usage

#### Simple
You can use it with the official driver, for example:

```go
package main

import (
	"context"
	"github.com/hfunc/hfunc-go/hfdatabase/hfmongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()
	db, _ := mongo.NewClient()
	col := db.Collection("test")
	builder := hfmongo.NewBuilder()
	builder.Field("name").Equal("nekilc")
	buidler.Field("age").LessThan(100)
	builder.Skip(0).Limit(2)
	col.Find(ctx,builder.Filter(),builder.FindOpts())
	buidler.Field("age").Set(200)
	col.Update(ctx,builder.Filter(), builder.Update(),builder.UpOpts())
}
```
MongoDB Find Statement:
```genericsql
db.test.find({"name":"nekilc","age":{$lt:100}},{},{skip:0,limit:2})
```

MongoDB Update Statement:
```mongo
db.test.update({"name":"nekilc","age":{$lt:100}},{$set:{age:200}})
```

You can use the package directly, for example:

```go

```

#### Logic Operator

```go
package main

import (
	"github.com/hfunc/hfunc-go/hfdatabase/hfmongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	db, _ := mongo.NewClient()
	col := db.Collection("test")
	builder := hfmongo.NewBuilder()
	// ------- Logic Start---------
	orBd := hfmongo.NewBuilder()
	orBd.Field("age").Equal(11)
	orBd.Field("name").Equal("hfunc")
	orBd2 := hfmongo.NewBuilder()
	orBd2.Field("name").Equal("hfunc1")
	// At the same level
	builder.Or(orBd).Or(orBd2)
	// At the different level
	builder.Or(orBd.Or(orBd2))
	// ------- Logic End---------
	buidler.Field("age").LessThan(100)
	col.Find(builder.Filter())
	buidler.Field("age").Set(200)
	col.Update(builder.Filter(),builder.Update())
}
```
You can replace the above Logic code block with the following code:
```go
// At the same level
builder.OrFc(func(bd *hfmongo.Builder){
	bd.Field("age").Equal(11)
    bd.Field("name").Equal("hfunc")
}).OrFc(func(bd *hfmongo.Builder){
    bd.Field("name").Equal("hfunc1")
})
// At the different level
builder.OrFc(func(bd *hfmongo.Builder){
    bd.Field("age").Equal(11)
    bd.Field("name").Equal("hfunc")
	bd.OrFc(func(bd *hfmongo.Builder){
        bd.Field("name").Equal("hfunc1")
    })
})
```
MongoDB Find Statement:
```genericsql
-- At the same level
db.test.Find({$or:[{age:11,name:"hfunc"},{name:"func1"}]})
-- At the different level
db.test.Find({$or:[{age:11,name:"hfunc",$or:[{name:"hfunc1"}]}]})
```