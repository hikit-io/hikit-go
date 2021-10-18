package hfmongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

func (f *Filed) Set(val Any) *Filed {
	return f.op(UpdateOp.Set, OpTypeUpdate, val)
}

func (f *Filed) UnSet() *Filed {
	return f.op(UpdateOp.Unset, OpTypeUpdate, "")
}

func (f *Filed) Push(val Any) *Filed {
	return f.op(UpdateOp.Push, OpTypeUpdate, val)
}

func (f *Filed) AddToSet(val Any) *Filed {
	if !IsAnys(val) {
		panic("val is array type")
	}
	return f.op(UpdateOp.AddToSet, OpTypeUpdate, val)
}
