package hkmongo

import (
	. "go.hikit.io/hktypes"
)

func (f *Field) Set(val Any) *Field {
	return f.op(UpdateOp.Set, OpTypeUpdate, val)
}

func (f *Field) UnSet() *Field {
	return f.op(UpdateOp.Unset, OpTypeUpdate, "")
}

func (f *Field) Push(val Any) *Field {
	return f.op(UpdateOp.Push, OpTypeUpdate, val)
}

func (f *Field) AddToSet(val Any) *Field {
	if !IsAnys(val) {
		panic("val is array type")
	}
	return f.op(UpdateOp.AddToSet, OpTypeUpdate, val)
}
