package hfmongo

import (
	"context"
	. "github.com/hfunc/hfunc-go/hftypes"
	"go.mongodb.org/mongo-driver/mongo"
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
		nameStr = n.DocName()
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

func (c *Collection) FindAny(ctx context.Context, val MustPtr, res MustSlicePtr) error {
	builder := &Builder{}
	switch inst := val.(type) {
	case map[string]interface{}:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]string:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int8:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int16:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int32:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int64:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint8:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint16:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint32:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint64:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case *Builder:
		builder = inst
	case Builder:
		builder = &inst
	default:
		rft := reflect.TypeOf(val).Elem()
		rfv := reflect.ValueOf(val).Elem()
		if rft.Kind() == reflect.Struct {
			for i := 0; i < rft.NumField(); i++ {
				if rfv.Field(i).Kind() == reflect.Ptr {
					if rfv.FieldByName(rft.Field(i).Name).IsZero() || rfv.FieldByName(rft.Field(i).Name).IsNil() {
						continue
					}
				}
				v, ok := rft.Field(i).Tag.Lookup("bson")
				if ok {
					builder.Field(v).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
					continue
				}

				builder.Field(rft.Field(i).Name).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
			}
		}
	}
	cur, err := c.Collection.Find(ctx, builder.Filter(), builder.FindOpts())
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

func (c *Collection) UpdateAny(ctx context.Context, val MustPtr, res MustSlicePtr) error {
	builder := &Builder{}
	switch inst := val.(type) {
	case map[string]interface{}:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]string:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int8:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int16:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int32:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]int64:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint8:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint16:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint32:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case map[string]uint64:
		for field, value := range inst {
			builder.Field(field).Equal(value)
		}
	case *Builder:
		builder = inst
	case Builder:
		builder = &inst
	default:
		rft := reflect.TypeOf(val).Elem()
		rfv := reflect.ValueOf(val).Elem()
		if rft.Kind() == reflect.Struct {
			for i := 0; i < rft.NumField(); i++ {
				if rfv.Field(i).Kind() == reflect.Ptr {
					if rfv.FieldByName(rft.Field(i).Name).IsZero() || rfv.FieldByName(rft.Field(i).Name).IsNil() {
						continue
					}
				}
				v, ok := rft.Field(i).Tag.Lookup("bson")
				if ok {
					builder.Field(v).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
					continue
				}

				builder.Field(rft.Field(i).Name).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
			}
		}
	}
	cur, err := c.Collection.Find(ctx, builder.Filter(), builder.FindOpts())
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

func NewDB(client *mongo.Client, dbname string) *Database {
	return &Database{
		dbname: dbname,
		Client: client,
		tables: map[TableName]*Collection{},
	}
}
