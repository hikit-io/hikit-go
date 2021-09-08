package mongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

func (f *Filed) Set(val Any) *Filed {
	return f.op(UpdateOp.Set, OpTypeUpdate, val)
}

func (f *Filed) UnSet() *Filed {
	return f.op(UpdateOp.Unset, OpTypeUpdate, "")
}
