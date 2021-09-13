package hfmongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

// In (hftypes.IsSlice(val) == true)
func (f *Filed) In(val Any) *Filed {
	return f.op(FindOp.In, OpTypeFind, val)
}

// NotIn (hftypes.IsSlice(val) == true)
func (f *Filed) NotIn(val Any) *Filed {
	return f.op(FindOp.NotIn, OpTypeFind, val)
}

func (f *Filed) Equal(val Any) *Filed {
	return f.op(FindOp.Equal, OpTypeFind, val)
}

func (f *Filed) NotEqual(val Any) *Filed {
	return f.op(FindOp.NotEqual, OpTypeFind, val)
}

func (f *Filed) GreatThan(val Any) *Filed {
	return f.op(FindOp.GreatThan, OpTypeFind, val)
}

func (f *Filed) GreatThanEqual(val Any) *Filed {
	return f.op(FindOp.GreatThanEqual, OpTypeFind, val)
}

func (f *Filed) LessThan(val Any) *Filed {
	return f.op(FindOp.LessThan, OpTypeFind, val)
}

func (f *Filed) LessThanEqual(val Any) *Filed {
	return f.op(FindOp.LessThanEqual, OpTypeFind, val)
}
