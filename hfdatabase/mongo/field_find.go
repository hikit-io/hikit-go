package mongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

func (f *Filed) In(val Any) *Filed {
	return f.op(FindOp.In, OpTypeFind, val)

}

func (f *Filed) Eq(val Any) *Filed {
	return f.op(FindOp.Equal, OpTypeFind, val)
}
