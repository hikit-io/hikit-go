package hkmg

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.hikit.io/hktypes"
)

type mergeOpts struct {
	f              *options.FindOptions
	u              *options.UpdateOptions
	returnDocument *options.ReturnDocument
}

func (m mergeOpts) ToFindOneAndUpdateOptions() *options.FindOneAndUpdateOptions {
	return &options.FindOneAndUpdateOptions{
		ArrayFilters:             m.u.ArrayFilters,
		BypassDocumentValidation: m.u.BypassDocumentValidation,
		Collation:                m.u.Collation,
		MaxTime:                  m.f.MaxTime,
		Projection:               m.f.Projection,
		ReturnDocument:           m.returnDocument,
		Sort:                     m.f.Sort,
		Upsert:                   m.u.Upsert,
		Hint:                     m.f.Hint,
	}
}

func (m mergeOpts) ToFindOneAndReplaceOptions() *options.FindOneAndReplaceOptions {
	return &options.FindOneAndReplaceOptions{
		BypassDocumentValidation: m.u.BypassDocumentValidation,
		Collation:                m.u.Collation,
		MaxTime:                  m.f.MaxTime,
		Projection:               m.f.Projection,
		ReturnDocument:           m.returnDocument,
		Sort:                     m.f.Sort,
		Upsert:                   m.u.Upsert,
		Hint:                     m.f.Hint,
	}
}

func (m mergeOpts) ToFindOneAndDeleteOptions() *options.FindOneAndDeleteOptions {
	return &options.FindOneAndDeleteOptions{
		Collation:  m.f.Collation,
		MaxTime:    m.f.MaxTime,
		Projection: m.f.Projection,
		Sort:       m.f.Sort,
		Hint:       m.f.Hint,
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

func (m mergeOpts) ToCountOptions() *options.CountOptions {
	return &options.CountOptions{
		Collation: m.f.Collation,
		Hint:      m.f.Hint,
		MaxTime:   m.f.MaxTime,
		Skip:      m.f.Skip,
		Limit:     m.f.Limit,
	}
}

func (m mergeOpts) ToDeleteOptions() *options.DeleteOptions {
	return &options.DeleteOptions{
		Collation: m.f.Collation,
		Hint:      m.f.Hint,
	}
}

func (m mergeOpts) ToReplaceOptions() *options.ReplaceOptions {
	return &options.ReplaceOptions{
		BypassDocumentValidation: m.u.BypassDocumentValidation,
		Collation:                m.f.Collation,
		Hint:                     m.f.Hint,
		Upsert:                   m.u.Upsert,
	}
}

func (c *Executor) Limit(i int64) *Executor {
	c.SetLimit(i)
	return c
}

func (c *Executor) Skip(i int64) *Executor {
	c.SetSkip(i)
	return c
}

func (c *Executor) Page(size, num int64) *Executor {
	c.Limit(size).Skip((num - 1) * size)
	return c
}

func (c *Executor) Upsert(enable bool) *Executor {
	c.SetUpsert(enable)
	return c
}

func (c *Executor) ShowRecordID(enable bool) *Executor {
	c.SetShowRecordID(enable)
	return c
}

func (c *Executor) BatchSize(i int32) *Executor {
	c.SetBatchSize(i)
	return c
}

func (c *Executor) BypassDocumentValidation(enable bool) *Executor {
	c.SetBypassDocumentValidation(enable)
	return c
}

func (c *Executor) Sort(sort hktypes.MustKV) *Executor {
	builder := NewBuilder().parseVal(sort, Sort, c.opt.fieldNameFc)
	c.SetSort(builder.FindOpts().Sort)
	return c
}

type ReturnDocType = options.ReturnDocument

const (
	Before ReturnDocType = options.Before
	After  ReturnDocType = options.After
)

func (c *Executor) ReturnDoc(t ReturnDocType) *Executor {
	c.returnDocument = &t
	return c
}

func (c *Executor) MaxTime(d time.Duration) *Executor {
	c.SetMaxTime(d)
	return c
}

func (c *Executor) AllowDiskUse(b bool) *Executor {
	c.SetAllowDiskUse(b)
	return c
}

func (c *Executor) AllowPartialResults(b bool) *Executor {
	c.SetAllowPartialResults(b)
	return c
}

func (c *Executor) MaxAwaitTime(d time.Duration) *Executor {
	c.SetMaxAwaitTime(d)
	return c
}

func (c *Executor) ReturnKey(b bool) *Executor {
	c.SetReturnKey(b)
	return c
}

type CursorType = options.CursorType

const (
	NonTailable   = options.NonTailable
	Tailable      = options.Tailable
	TailableAwait = options.TailableAwait
)

func (c *Executor) CursorType(ct CursorType) *Executor {
	c.SetCursorType(ct)
	return c
}

func (c *Executor) Comment(comment string) *Executor {
	c.SetComment(comment)
	return c
}

func (c *Executor) NoCursorTimeout(b bool) *Executor {
	c.SetNoCursorTimeout(b)
	return c
}
