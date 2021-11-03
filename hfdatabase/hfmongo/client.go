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

func (c *Collection) HInsertOne(ctx context.Context, doc Any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, Error) {
	r, e := c.Collection.InsertOne(ctx, doc, opts...)
	return r, err{e}
}

func (c *Collection) HInsertMany(ctx context.Context, docs Any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, Error) {
	var idocs Anys
	switch i := docs.(type) {
	case []interface{}:
		idocs = i
	default:
		idocs = AnyToSliceAny(docs)
	}
	r, e := c.Collection.InsertMany(ctx, idocs, opts...)
	return r, err{e}
}

func (c *Collection) HFindOne(ctx context.Context, val MustPtr, res MustSlicePtr, opts ...*options.FindOneOptions) (*mongo.SingleResult, Error) {
	builder := NewBuilder().parseVal(val, Find)
	opt := options.MergeFindOneOptions(append(opts, mergeOpts{f: builder.FindOpts()}.ToFindOneOptions())...)
	r := c.Collection.FindOne(ctx, builder.Filter(), opt)
	if r.Err() != nil {
		return r, err{
			r.Err(),
		}
	}
	e := r.Decode(res)
	if e != nil {
		return r, err{e}
	}
	return r, nil
}

type mergeOpts struct {
	f *options.FindOptions
	u *options.UpdateOptions
}

func (m mergeOpts) ToFindOneAndUpdateOptions() *options.FindOneAndUpdateOptions {
	return &options.FindOneAndUpdateOptions{
		ArrayFilters:             m.u.ArrayFilters,
		BypassDocumentValidation: m.u.BypassDocumentValidation,
		Collation:                m.u.Collation,
		MaxTime:                  m.f.MaxTime,
		Projection:               m.f.Projection,
		//ReturnDocument:    todo       ,
		Sort:   m.f.Sort,
		Upsert: m.u.Upsert,
		Hint:   m.f.Hint,
	}
}

func (m mergeOpts) ToFindOneAndReplaceOptions() *options.FindOneAndReplaceOptions {
	return &options.FindOneAndReplaceOptions{
		BypassDocumentValidation: m.u.BypassDocumentValidation,
		Collation:                m.u.Collation,
		MaxTime:                  m.f.MaxTime,
		Projection:               m.f.Projection,
		//ReturnDocument:    todo       ,
		Sort:   m.f.Sort,
		Upsert: m.u.Upsert,
		Hint:   m.f.Hint,
	}
}

func (m mergeOpts) ToFindOneOptions() *options.FindOneOptions {
	return &options.FindOneOptions{
		AllowPartialResults: m.f.AllowPartialResults,
		BatchSize:           m.f.BatchSize,
		Collation:           m.f.Collation,
		Comment:             m.f.Comment,
		CursorType:          m.f.CursorType,
		Hint:                m.f.Hint,
		Max:                 m.f.Max,
		MaxAwaitTime:        m.f.MaxAwaitTime,
		MaxTime:             m.f.MaxTime,
		Min:                 m.f.Min,
		NoCursorTimeout:     m.f.NoCursorTimeout,
		OplogReplay:         m.f.OplogReplay,
		Projection:          m.f.Projection,
		ReturnKey:           m.f.ReturnKey,
		ShowRecordID:        m.f.ShowRecordID,
		Skip:                m.f.Skip,
		Snapshot:            m.f.Snapshot,
		Sort:                m.f.Sort,
	}
}

func (c *Collection) HFindOneAndUpdate(ctx context.Context, condition MustPtr, update, updateRes MustSlicePtr, opts ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, Error) {
	builder := NewBuilder().parseVal(condition, Find).parseVal(update, Update).parseVal(updateRes, Projection)
	opt := options.MergeFindOneAndUpdateOptions(append(opts, mergeOpts{builder.FindOpts(), builder.UpOpts()}.ToFindOneAndUpdateOptions())...)
	r := c.Collection.FindOneAndUpdate(ctx,
		builder.Filter(), builder.Update(),
		opt,
	)
	if r.Err() != nil {
		return r, err{
			r.Err(),
		}
	}
	return r, err{
		r.Decode(updateRes),
	}
}

func (c *Collection) HFindOneAndReplace(ctx context.Context, condition, replace Any, res MustSlicePtr, opts ...*options.FindOneAndReplaceOptions) (*mongo.SingleResult, Error) {
	builder := NewBuilder().parseVal(condition, Find).parseVal(res, Projection)
	opt := options.MergeFindOneAndReplaceOptions(append(opts, mergeOpts{builder.FindOpts(), builder.UpOpts()}.ToFindOneAndReplaceOptions())...)

	r := c.Collection.FindOneAndReplace(ctx,
		builder.Filter(), replace,
		opt)
	if r.Err() != nil {
		return r, err{
			r.Err(),
		}
	}
	return r, err{
		r.Decode(res),
	}
}

type parseType uint8

const (
	Find parseType = iota + 1
	Update
	Projection
)

func (b *Builder) parseVal(val MustPtr, pt parseType) *Builder {
	//b := NewBuilder()
	switch inst := val.(type) {
	case map[string]interface{}:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]string:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]int:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]int8:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]int16:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]int32:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]int64:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]uint:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]uint8:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]uint16:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]uint32:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case map[string]uint64:
		for field, value := range inst {
			switch pt {
			case Find:
				b.Field(field).Equal(value)
			case Update:
				b.Field(field).Set(value)
			case Projection:
				b.Field(field).Projection(true)
			}
		}
	case *Builder:
		*b = *inst
	case Builder:
		*b = inst
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
				if rfv.FieldByName(rft.Field(i).Name).IsZero() {
					continue
				}
				v, ok := rft.Field(i).Tag.Lookup("bson")
				if ok {
					switch pt {
					case Find:
						b.Field(v).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
					case Update:
						b.Field(v).Set(rfv.FieldByName(rft.Field(i).Name).Interface())

					case Projection:
						b.Field(v).Projection(true)
					}
					continue
				}
				switch pt {
				case Find:
					b.Field(rft.Field(i).Name).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
				case Update:
					b.Field(rft.Field(i).Name).Set(rfv.FieldByName(rft.Field(i).Name).Interface())
				case Projection:
					b.Field(rft.Field(i).Name).Projection(true)
				}
			}
		}
	}
	return b
}

func (c *Collection) HFind(ctx context.Context, condition MustPtr, res MustSlicePtr, opts ...*options.FindOptions) Error {
	builder := NewBuilder().parseVal(condition, Find).parseVal(res, Projection)
	opt := options.MergeFindOptions(append(opts, builder.FindOpts())...)

	cur, e := c.Collection.Find(ctx, builder.Filter(), opt)
	if e != nil {
		return err{e}
	}
	if e = cur.All(ctx, res); e != nil {
		return err{e}
	}
	return nil
}

func (c *Collection) HUpdateOne(ctx context.Context, condition, update MustPtr, opts ...*options.UpdateOptions) (*mongo.UpdateResult, Error) {
	builder := NewBuilder().parseVal(condition, Find).parseVal(update, Update)
	opt := options.MergeUpdateOptions(append(opts, builder.UpOpts())...)

	cur, e := c.Collection.UpdateOne(ctx, builder.Filter(), builder.Update(), opt)
	return cur, err{e}
}

func (c *Collection) HUpdateMany(ctx context.Context, condition, update MustPtr, opts ...*options.UpdateOptions) (*mongo.UpdateResult, Error) {
	builder := NewBuilder().parseVal(condition, Find).parseVal(update, Update)
	opt := options.MergeUpdateOptions(append(opts, builder.UpOpts())...)

	cur, e := c.Collection.UpdateMany(ctx, builder.Filter(), builder.Update(), opt)
	return cur, err{e}
}

func (c *Collection) Count(ctx context.Context) {
	//c.Collection.CountDocuments()
	//todo
}

func (c *Collection) HDeleteOne() {
	//c.Collection.DeleteOne()
	// todo
}

func (c *Collection) HDeleteMany() {
	//c.Collection.DeleteOne()
	// todo
}

func (c *Collection) HFindOneAndDelete() {
	//c.Collection.FindOneAndDelete()
	// todo
}

func NewDB(client *mongo.Client, dbname string) *Database {
	return &Database{
		dbname: dbname,
		Client: client,
		tables: map[TableName]*Collection{},
	}
}
