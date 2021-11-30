package hkmg

import (
	"context"
	"reflect"

	. "go.hikit.io/hktypes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Executor struct {
	opt    *Options
	parent *Collection
	*options.FindOptions
	*options.UpdateOptions
}

func (c *Executor) HInsertOne(ctx context.Context, doc MustKV, opts ...*options.InsertOneOptions) *InsertOneResult {
	r, e := c.parent.InsertOne(ctx, doc, opts...)
	return &InsertOneResult{err{e}, r}
}

func (c *Executor) HInsertMany(ctx context.Context, docs MustKV, opts ...*options.InsertManyOptions) *InsertManyResult {
	var idocs Anys
	switch i := docs.(type) {
	case []interface{}:
		idocs = i
	default:
		idocs = AnyToSliceAny(docs)
	}
	r, e := c.parent.InsertMany(ctx, idocs, opts...)
	return &InsertManyResult{err{e}, r}
}

func (c *Executor) HFindOne(ctx context.Context, val MustKV, res MustPtr, opts ...*options.FindOneOptions) *SingleResult {
	builder := NewBuilder().parseVal(val, Find, c.opt.fieldNameFc)
	opt := options.MergeFindOneOptions(append(opts,
		mergeOpts{f: builder.FindOpts()}.ToFindOneOptions(),
		mergeOpts{f: c.FindOptions}.ToFindOneOptions(),
	)...)
	r := c.parent.FindOne(ctx, builder.Filter(), opt)
	if r.Err() != nil {
		return &SingleResult{
			err{r.Err()},
		}
	}
	return &SingleResult{
		err{r.Decode(res)},
	}
}

func (c *Executor) HFindOneAndUpdate(ctx context.Context, condition, update MustKV, updateRes MustPtr, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(update, Update, c.opt.fieldNameFc).parseVal(updateRes, Projection, c.opt.fieldNameFc)
	opt := options.MergeFindOneAndUpdateOptions(append(opts,
		mergeOpts{builder.FindOpts(), builder.UpOpts()}.ToFindOneAndUpdateOptions(),
		mergeOpts{f: c.FindOptions, u: c.UpdateOptions}.ToFindOneAndUpdateOptions(),
	)...)
	r := c.parent.FindOneAndUpdate(ctx,
		builder.Filter(), builder.Update(),
		opt,
	)
	if r.Err() != nil {
		return &SingleResult{
			err: err{r.Err()},
		}
	}
	return &SingleResult{
		err: err{r.Decode(updateRes)},
	}
}

func (c *Executor) HFindOneAndReplace(ctx context.Context, condition, replace MustKV, res MustPtr, opts ...*options.FindOneAndReplaceOptions) *SingleResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(res, Projection, c.opt.fieldNameFc)
	opt := options.MergeFindOneAndReplaceOptions(append(opts,
		mergeOpts{builder.FindOpts(), builder.UpOpts()}.ToFindOneAndReplaceOptions(),
		mergeOpts{c.FindOptions, c.UpdateOptions}.ToFindOneAndReplaceOptions(),
	)...)

	r := c.parent.FindOneAndReplace(ctx,
		builder.Filter(), replace,
		opt)
	if r.Err() != nil {
		return &SingleResult{
			err: err{r.Err()},
		}
	}
	return &SingleResult{
		err: err{r.Decode(res)},
	}
}

func (c *Executor) HFindOneAndDelete(ctx context.Context, condition MustKV, deleteRes MustPtr, opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(deleteRes, Projection, c.opt.fieldNameFc)
	opt := options.MergeFindOneAndDeleteOptions(append(opts,
		mergeOpts{f: builder.FindOpts()}.ToFindOneAndDeleteOptions(),
		mergeOpts{f: c.FindOptions}.ToFindOneAndDeleteOptions(),
	)...)
	r := c.parent.FindOneAndDelete(ctx,
		builder.Filter(),
		opt,
	)
	if r.Err() != nil {
		return &SingleResult{
			err: err{r.Err()},
		}
	}
	return &SingleResult{
		err: err{r.Decode(deleteRes)},
	}
}

type parseType uint8

const (
	Find parseType = iota + 1
	Update
	Projection
)

func (b *Builder) parseVal(val MustKV, pt parseType, format FieldNameFormat) *Builder {
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
		rft := reflect.TypeOf(val)
		rfv := reflect.ValueOf(val)
		if rft.Kind() == reflect.Ptr {
			rft = rft.Elem()
			rfv = rfv.Elem()
		}
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
					if format != nil {
						b.Field(format(rft.Field(i).Name)).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
					} else {
						b.Field(rft.Field(i).Name).Equal(rfv.FieldByName(rft.Field(i).Name).Interface())
					}
				case Update:
					if format != nil {
						b.Field(format(rft.Field(i).Name)).Set(rfv.FieldByName(rft.Field(i).Name).Interface())
					} else {
						b.Field(rft.Field(i).Name).Set(rfv.FieldByName(rft.Field(i).Name).Interface())
					}
				case Projection:
					if format != nil {
						b.Field(format(rft.Field(i).Name)).Projection(true)
					} else {
						b.Field(rft.Field(i).Name).Projection(true)
					}
				}
			}
		}
	}
	return b
}

func (c *Executor) HFind(ctx context.Context, condition MustKV, res MustSlicePtr, opts ...*options.FindOptions) *FindResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(res, Projection, c.opt.fieldNameFc)
	opt := options.MergeFindOptions(append(opts,
		builder.FindOpts(),
		c.FindOptions,
	)...)

	cur, e := c.parent.Find(ctx, builder.Filter(), opt)
	if e != nil {
		return &FindResult{
			err{e},
		}
	}
	return &FindResult{err{cur.All(ctx, res)}}
}

func (c *Executor) HUpdateOne(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(update, Update, c.opt.fieldNameFc)
	opt := options.MergeUpdateOptions(append(opts,
		builder.UpOpts(),
		c.UpdateOptions,
	)...)

	cur, e := c.parent.UpdateOne(ctx, builder.Filter(), builder.Update(), opt)
	return &UpdateResult{
		err:          err{e},
		UpdateResult: cur,
	}
}

func (c *Executor) HUpdateMany(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc).parseVal(update, Update, c.opt.fieldNameFc)
	opt := options.MergeUpdateOptions(append(opts,
		builder.UpOpts(),
		c.UpdateOptions,
	)...)

	cur, e := c.parent.UpdateMany(ctx, builder.Filter(), builder.Update(), opt)
	return &UpdateResult{
		err:          err{e},
		UpdateResult: cur,
	}
}

func (c *Executor) HCount(ctx context.Context, condition MustKV, opts ...*options.CountOptions) *CountResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc)
	opt := options.MergeCountOptions(append(opts,
		mergeOpts{f: builder.FindOpts()}.ToCountOptions(),
		mergeOpts{f: c.FindOptions}.ToCountOptions(),
	)...)
	count, e := c.parent.CountDocuments(ctx, builder.Filter(), opt)
	return &CountResult{
		err:   err{e},
		Count: count,
	}
}

func (c *Executor) HDeleteOne(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc)
	opt := options.MergeDeleteOptions(append(opts,
		mergeOpts{f: builder.FindOpts()}.ToDeleteOptions(),
		mergeOpts{f: c.FindOptions}.ToDeleteOptions(),
	)...)
	r, e := c.parent.DeleteOne(ctx, builder.Filter(), opt)
	return &DeleteResult{
		err:          err{e},
		DeleteResult: r,
	}
}

func (c *Executor) HDeleteMany(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc)
	opt := options.MergeDeleteOptions(append(opts,
		mergeOpts{f: builder.FindOpts()}.ToDeleteOptions(),
		mergeOpts{f: c.FindOptions}.ToDeleteOptions(),
	)...)
	r, e := c.parent.DeleteMany(ctx, builder.Filter(), opt)
	return &DeleteResult{
		err:          err{e},
		DeleteResult: r,
	}
}

func (c *Executor) HReplaceOne(ctx context.Context, condition, newDoc MustKV, opts ...*options.ReplaceOptions) *UpdateResult {
	builder := NewBuilder().parseVal(condition, Find, c.opt.fieldNameFc)
	opt := options.MergeReplaceOptions(append(opts,
		mergeOpts{f: builder.FindOpts(), u: builder.UpOpts()}.ToReplaceOptions(),
		mergeOpts{f: c.FindOptions, u: c.UpdateOptions}.ToReplaceOptions(),
	)...)
	r, e := c.parent.ReplaceOne(ctx, builder.Filter(), newDoc, opt)
	return &UpdateResult{
		err:          err{e},
		UpdateResult: r,
	}
}
