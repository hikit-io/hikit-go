package hfmongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

func (f *Filed) LessThan(val Any) *Filed {
	return f.op(FindOp.LessThan, OpTypeFind, val)
}

func (f *Filed) LessThanEqual(val Any) *Filed {
	return f.op(FindOp.LessThanEqual, OpTypeFind, val)
}

func (f *Filed) GreatThan(val Any) *Filed {
	return f.op(FindOp.GreatThan, OpTypeFind, val)
}

func (f *Filed) GreatThanEqual(val Any) *Filed {
	return f.op(FindOp.GreatThanEqual, OpTypeFind, val)
}

func (f *Filed) NotEqual(val Any) *Filed {
	return f.op(FindOp.NotEqual, OpTypeFind, val)
}

func (f *Filed) All(val Any) *Filed {
	return f.op(FindOp.All, OpTypeFind, val)
}

// In (hftypes.IsSlice(val) == true)
func (f *Filed) In(val Any) *Filed {
	return f.op(FindOp.In, OpTypeFind, val)
}

// NotIn (hftypes.IsSlice(val) == true)
func (f *Filed) NotIn(val Any) *Filed {
	return f.op(FindOp.NotIn, OpTypeFind, val)
}

func (f *Filed) Size(val Ui) *Filed {
	return f.op(FindOp.Size, OpTypeFind, val)
}

func (f *Filed) Exists(val Any) *Filed {
	return f.op(FindOp.Exists, OpTypeFind, val)
}

func (f *Filed) Type(val Any) *Filed {
	return f.op(FindOp.Type, OpTypeFind, val)
}

func (f *Filed) Mod(val Any) *Filed {
	return f.op(FindOp.Mod, OpTypeFind, val)
}

func (f *Filed) Regex(val Str) *Filed {
	return f.op(FindOp.Regex, OpTypeFind, val)
}

func (f *Filed) Text(val Any) *Filed {
	return f.op(FindOp.Text, OpTypeFind, val)
}

func (f *Filed) ElemMatch(val MustStructOrPtr) *Filed {
	return f.op(FindOp.ElemMatch, OpTypeFind, val)
}

func (f *Filed) Equal(val Any) *Filed {
	return f.op(FindOp.Equal, OpTypeFind, val)
}

func (f *Filed) Hint(val int) *Filed {
	return f.op(QueryOp.Hint, OpTypeQuery, val)
}

type SortType int

const (
	//SortDesc 降序
	SortDesc SortType = -1 + iota
	SortNone
	//SortAsc 升序
	SortAsc
)

func (f *Filed) Sort(val SortType) *Filed {
	return f.op(QueryOp.Sort, OpTypeQuery, val)
}

func (f *Filed) Projection(val bool) *Filed {
	return f.op(QueryOp.Projection, OpTypeQuery, val)
}
