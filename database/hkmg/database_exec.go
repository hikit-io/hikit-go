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
func (c *DbExecutor) HInsertOne(ctx context.Context, doc Any, opts ...*options.InsertOneOptions) *InsertOneResult {
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

func (c *DbExecutor) HInsertMany(ctx context.Context, docs Any, opts ...*options.InsertManyOptions) *InsertManyResult {
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

type DbExecutor struct {
	f *options.FindOptions
	u *options.UpdateOptions
	*Database
}

func (c *Database) Exec() *DbExecutor {
	return &DbExecutor{
		f:        options.Find(),
		u:        options.Update(),
		Database: c,
	}
}

func (c *DbExecutor) HFindOne(ctx context.Context, condition MustKV, result MustPtr, opts ...*options.FindOneOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).HFindOne(ctx, condition, result, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).HFindOne(ctx, cond, result, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).HFindOne(ctx, condition, result, opts...)
	}
}

func (c *DbExecutor) HFindOneAndUpdate(ctx context.Context, condition, update MustKV, updateRes MustPtr, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndUpdate(ctx, condition, update, updateRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndUpdate(ctx, cond, update, updateRes, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndUpdate(ctx, condition, update, updateRes, opts...)
	}
}

func (c *DbExecutor) HFindOneAndReplace(ctx context.Context, condition, replace MustKV, replaceRes MustPtr, opts ...*options.FindOneAndReplaceOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndReplace(ctx, condition, replace, replaceRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndReplace(ctx, cond, replace, replaceRes, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndReplace(ctx, condition, replace, replaceRes, opts...)
	}
}

func (c *DbExecutor) HFindOneAndDelete(ctx context.Context, condition MustKV, updateRes MustPtr, opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndDelete(ctx, condition, updateRes, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndDelete(ctx, cond, updateRes, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFindOneAndDelete(ctx, condition, updateRes, opts...)
	}
}

func (c *DbExecutor) HFind(ctx context.Context, condition MustKV, res MustSlicePtr, opts ...*options.FindOptions) *FindResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFind(ctx, condition, res, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFind(ctx, cond, res, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HFind(ctx, condition, res, opts...)
	}
}

func (c *DbExecutor) HUpdateOne(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateOne(ctx, condition, update, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateOne(ctx, cond, update, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateOne(ctx, condition, update, opts...)
	}
}

func (c *DbExecutor) HUpdateMany(ctx context.Context, condition, update MustKV, opts ...*options.UpdateOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateMany(ctx, condition, update, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateMany(ctx, cond, update, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HUpdateMany(ctx, condition, update, opts...)
	}
}

func (c *DbExecutor) HCount(ctx context.Context, condition MustKV, opts ...*options.CountOptions) *CountResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HCount(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HCount(ctx, cond, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HCount(ctx, condition, opts...)
	}
}

func (c *DbExecutor) HDeleteOne(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteOne(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteOne(ctx, cond, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteOne(ctx, condition, opts...)
	}
}

func (c *DbExecutor) HDeleteMany(ctx context.Context, condition MustKV, opts ...*options.DeleteOptions) *DeleteResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteMany(ctx, condition, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteMany(ctx, cond, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HDeleteMany(ctx, condition, opts...)
	}
}

func (c *DbExecutor) HReplaceOne(ctx context.Context, condition, newDoc MustKV, opts ...*options.ReplaceOptions) *UpdateResult {
	switch d := condition.(type) {
	case Doc:
		return c.Col(d.DocName()).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HReplaceOne(ctx, condition, newDoc, opts...)
	case QueryFc:
		cond, table := d()
		return c.Col(table).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HReplaceOne(ctx, cond, newDoc, opts...)
	default:
		return c.Col(condition).SetFindOptions(*c.f).SetUpdateOptions(*c.u).HReplaceOne(ctx, condition, newDoc, opts...)
	}
}
