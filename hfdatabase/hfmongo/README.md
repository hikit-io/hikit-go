# hfmongo

## Features

- Chain API
- Custom log

## Getting start

`go get github.com/hfunc/hfunc-go/hfdatabase/hfmongo`

### Usage

You can use it with the official driver, for example:

```go
package main

import (
	"github.com/hfunc/hfunc-go/hfdatabase/hfmongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	db, _ := mongo.NewClient()
	col := db.Collection("test")
	builder := &hfmongo.Builder{}
	builder.Field("name").Equal("nekilc")
	buidler.Field("age").LessThan(100)
	col.Find(builder.Filter())
	buidler.Field("age").Set(200)
	col.Update(builder.Filter(),builder.Update())
}
```
MongoDB Find Statement:
```genericsql
db.test.find({"name":"nekilc","age":{$lt:100}})
```

MongoDB Update Statement:
```mongo
db.test.update({"name":"nekilc","age":{$lt:100}},{$set:{age:200}})
```

You can use the package directly, for example:

```go

```


