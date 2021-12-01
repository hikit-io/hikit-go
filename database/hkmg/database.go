package hkmg

import (
	"context"
	"reflect"

	. "go.hikit.io/hktypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TableName = string

type TableNameFormat func(name string) string
type FieldNameFormat func(name string) string

type Options struct {
	tableNameFc TableNameFormat
	fieldNameFc FieldNameFormat
	debug       bool
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

func (c *Database) Col(model Any) *Executor {
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
	return &Executor{
		parent:        col,
		opt:           &c.options,
		FindOptions:   options.Find(),
		UpdateOptions: options.Update(),
	}
}

func NewDB(client *mongo.Client, dbname string, opts ...Option) *Database {
	opt := &Options{}
	for _, o := range opts {
		o.apply(opt)
	}
	db := &Database{
		options: *opt,
		dbname:  dbname,
		Client:  client,
		tables:  map[TableName]*Collection{},
	}
	tableNames, _ := db.DB().ListCollectionNames(context.Background(), bson.D{})
	for _, name := range tableNames {
		db.Col(name)
	}
	return db
}
