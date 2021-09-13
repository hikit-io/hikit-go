package hfmongo

import (
	"context"
	. "github.com/hfunc/hfunc-go/hftypes"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type Database struct {
	dbname string
	*mongo.Client
}

func (c *Database) DB() *mongo.Database {
	return c.Database(c.dbname)
}

type Collection struct {
	name Str
	*mongo.Collection
}

func (c *Database) Col(name Any) *Collection {
	nameStr := ""
	switch n := name.(type) {
	case string:
		nameStr = n
	case Doc:
		nameStr = n.Name()
	}
	return &Collection{
		nameStr,
		c.DB().Collection(nameStr),
	}
}

func (c *Collection) FindAny(ctx context.Context, val Any, res Any) error {
	rft := reflect.TypeOf(val).Elem()
	rfv := reflect.ValueOf(val).Elem()
	builder := Builder{}
	switch rft.Kind() {
	case reflect.Struct:
		for i := 0; i < rft.NumField(); i++ {
			v, ok := rft.Field(i).Tag.Lookup("json")
			if ok {
				builder.Field(v).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
				continue
			}
			builder.Field(rft.Field(i).Name).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
		}
	}

	cur, err := c.Find(ctx, builder.BuildFilter())
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

func (c *Collection) UpdateAny() {

}
