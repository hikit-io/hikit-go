package hkmg

import (
	"context"

	. "go.hikit.io/hktypes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryFc func() (condition interface{}, table interface{})

// HInsertOne
// Usage:
//
//		type table struct {
//			Name string `bson:"name"`
//			Age uint8 `bson:"age"`
//			Addr string `bson:"addr"`
//		}
//		type needTable struct {
//			Name string `bson:"name"`
//			Age uint8 `bson:"age"`
//		}
//		db.HInsertOne(ctx,table{})
//		db.HInsertOne(ctx,func()(interface{},interface{}){
//			return needTable{},table{}
//		})
//		db.HInsertOne(ctx,func()(interface{},interface{}){
//			return needTable{},"table"
//		})
func (c *Database) HInsertOne(ctx context.Context, doc Any, opts ...*options.InsertOneOptions) *InsertOneResult {
	switch d := doc.(type) {
	case Doc:
		return c.Col(d.DocName()).HInsertOne(ctx, doc, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HInsertOne(ctx, cond, opts...)
	default:
		return c.Col(doc).HInsertOne(ctx, doc, opts...)
	}
}

func (c *Database) HInsertMany(ctx context.Context, docs Any, opts ...*options.InsertManyOptions) *InsertManyResult {
	switch d := docs.(type) {
	case Doc:
		return c.Col(d.DocName()).HInsertMany(ctx, docs, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HInsertMany(ctx, cond, opts...)
	default:
		return c.Col(docs).HInsertMany(ctx, docs, opts...)
	}
}

func (c *Database) HFindOne(ctx context.Context, condition MustKV, result MustPtr, opts ...*options.FindOneOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HFindOne(ctx, condition, result, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HFindOne(ctx, cond, result, opts...)
	default:
		return c.Col(result).HFindOne(ctx, condition, result, opts...)
	}
}

func (c *Database) HFindOneAndUpdate(ctx context.Context, condition, update MustKV, updateRes MustPtr, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HFindOneAndUpdate(ctx, condition, update, updateRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HFindOneAndUpdate(ctx, cond, update, updateRes, opts...)
	default:
		return c.Col(updateRes).HFindOneAndUpdate(ctx, condition, update, updateRes, opts...)
	}
}

func (c *Database) HFindOneAndReplace(ctx context.Context, condition, replace MustKV, replaceRes MustPtr, opts ...*options.FindOneAndReplaceOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HFindOneAndReplace(ctx, condition, replace, replaceRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HFindOneAndReplace(ctx, cond, replace, replaceRes, opts...)
	default:
		return c.Col(replace).HFindOneAndReplace(ctx, condition, replace, replaceRes, opts...)
	}
}

func (c *Database) HFindOneAndDelete(ctx context.Context, condition MustKV, updateRes MustPtr, opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HFindOneAndDelete(ctx, condition, updateRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HFindOneAndDelete(ctx, cond, updateRes, opts...)
	default:
		return c.Col(condition).HFindOneAndDelete(ctx, condition, updateRes, opts...)
	}
}

func (c *Database) HFind(ctx context.Context, condition MustKV, res MustSlicePtr, opts ...*options.FindOptions) *FindResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HFind(ctx, condition, res, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HFind(ctx, cond, res, opts...)
	default:
		return c.Col(condition).HFind(ctx, condition, res, opts...)
	}
}

func (c *Database) HUpdateOne(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HUpdateOne(ctx, condition, update, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HUpdateOne(ctx, cond, update, opts...)
	default:
		return c.Col(condition).HUpdateOne(ctx, condition, update, opts...)
	}
}

func (c *Database) HUpdateMany(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HUpdateMany(ctx, condition, update, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HUpdateMany(ctx, cond, update, opts...)
	default:
		return c.Col(condition).HUpdateMany(ctx, condition, update, opts...)
	}
}

func (c *Database) HCount(ctx context.Context, condition MustKV, opts ...*options.CountOptions) *CountResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HCount(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HCount(ctx, cond, opts...)
	default:
		return c.Col(condition).HCount(ctx, condition, opts...)
	}
}

func (c *Database) HDeleteOne(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HDeleteOne(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HDeleteOne(ctx, cond, opts...)
	default:
		return c.Col(condition).HDeleteOne(ctx, condition, opts...)
	}
}

func (c *Database) HDeleteMany(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HDeleteMany(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HDeleteMany(ctx, cond, opts...)
	default:
		return c.Col(condition).HDeleteMany(ctx, condition, opts...)
	}
}

func (c *Database) HReplaceOne(ctx context.Context, condition, newDoc MustKV, opts ...*options.ReplaceOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).HReplaceOne(ctx, condition, newDoc, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).HReplaceOne(ctx, cond, newDoc, opts...)
	default:
		return c.Col(newDoc).HReplaceOne(ctx, condition, newDoc, opts...)
	}
}
