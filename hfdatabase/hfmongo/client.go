package hfmongo

import (
	"context"
	. "github.com/hfunc/hfunc-go/hftypes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type TableName = string

type TableNameFormat func(name string) string

type Options struct {
	tableNameFc func(s string) string
}

type Database struct {
	dbname string
	*mongo.Client
	tables  map[TableName]*Collection
	options Options
}

func (c *Database) DB() *mongo.Database {
	return c.Database(c.dbname)
}

type Collection struct {
	name Str
	*mongo.Collection
}

func (c *Database) Col(model Any) *Collection {
	nameStr := ""
	switch n := model.(type) {
	case string:
		nameStr = n
	case Doc:
		nameStr = n.Name()
	default:
		rv := reflect.ValueOf(model)
		if rv.Kind() == reflect.Ptr {
			nameStr = rv.Type().Elem().Name()
		}
		if rv.Kind() == reflect.Struct {
			nameStr = rv.Type().Name()
		}
	}
	if c.options.tableNameFc != nil {
		nameStr = c.options.tableNameFc(nameStr)
	} else {
		nameStr = ""
	}
	col, ok := c.tables[nameStr]
	if !ok {
		col = &Collection{
			name:       nameStr,
			Collection: c.DB().Collection(nameStr),
		}
		c.tables[nameStr] = col
	}
	return col
}

func (c *Collection) FindAny(ctx context.Context, val MustStructPtr, res MustSlicePtr) error {
	builder := Builder{}
	switch inst := val.(type) {
	case map[string]interface{}:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	}

	rft := reflect.TypeOf(val).Elem()
	rfv := reflect.ValueOf(val).Elem()
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
	options.Find()
	cur, err := c.Collection.Find(ctx, builder.Filter())
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

func (c *Collection) UpdateAny() {

}

func New(client *mongo.Client, daname string) *Database {
	return &Database{
		dbname: daname,
		Client: client,
		tables: map[TableName]*Collection{},
	}
}
