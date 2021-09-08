package mongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
)

var LogicOp = _LogicOp{
	And: "$and",
	Not: "$not",
	Nor: "$nor",
	Or:  "$or",
}

type _LogicOp struct {
	And Str
	Not Str
	Nor Str
	Or  Str
}
