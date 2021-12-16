package hkmg

import (
	"time"

	"go.hikit.io/hktypes"
)

func (c *DbExecutor) Limit(i int64) *DbExecutor {
	c.f.SetLimit(i)
	return c
}

func (c *DbExecutor) Skip(i int64) *DbExecutor {
	c.f.SetSkip(i)
	return c
}

func (c *DbExecutor) Page(size, num int64) *DbExecutor {
	c.Limit(size).Skip((num - 1) * size)
	return c
}

func (c *DbExecutor) Upsert(enable bool) *DbExecutor {
	c.u.SetUpsert(enable)
	return c
}

func (c *DbExecutor) ShowRecordID(enable bool) *DbExecutor {
	c.f.SetShowRecordID(enable)
	return c
}

func (c *DbExecutor) BatchSize(i int32) *DbExecutor {
	c.f.SetBatchSize(i)
	return c
}

func (c *DbExecutor) BypassDocumentValidation(enable bool) *DbExecutor {
	c.u.SetBypassDocumentValidation(enable)
	return c
}

func (c *DbExecutor) Sort(sort hktypes.MustKV) *DbExecutor {
	builder := NewBuilder().parseVal(sort, Sort, c.opt.fieldNameFc)
	c.f.SetSort(builder.FindOpts().Sort)
	return c
}

func (c *DbExecutor) ReturnDoc(t ReturnDocType) *DbExecutor {
	c.returnDocument = &t
	return c
}

func (c *DbExecutor) MaxTime(d time.Duration) *DbExecutor {
	c.f.SetMaxTime(d)
	return c
}

func (c *DbExecutor) AllowDiskUse(b bool) *DbExecutor {
	c.f.SetAllowDiskUse(b)
	return c
}

func (c *DbExecutor) AllowPartialResults(b bool) *DbExecutor {
	c.f.SetAllowPartialResults(b)
	return c
}

func (c *DbExecutor) MaxAwaitTime(d time.Duration) *DbExecutor {
	c.f.SetMaxAwaitTime(d)
	return c
}

func (c *DbExecutor) ReturnKey(b bool) *DbExecutor {
	c.f.SetReturnKey(b)
	return c
}

func (c *DbExecutor) CursorType(ct CursorType) *DbExecutor {
	c.f.SetCursorType(ct)
	return c
}

func (c *DbExecutor) Comment(comment string) *DbExecutor {
	c.f.SetComment(comment)
	return c
}

func (c *DbExecutor) NoCursorTimeout(b bool) *DbExecutor {
	c.f.SetNoCursorTimeout(b)
	return c
}
