package hkmongo

import (
	. "go.hikit.io/hikit/hktypes"
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
