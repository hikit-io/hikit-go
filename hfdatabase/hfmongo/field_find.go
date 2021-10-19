package hfmongo

import (
	. "github.com/hfunc/hfunc-go/hftypes"
	"regexp"
	"strconv"
	"strings"
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

func (f *Filed) Expression(expr string) *Filed {
	expr = strings.ReplaceAll(expr, " ", "")
	e, _ := regexp.Compile(`^[0-9]+((<)|(<=)|(>)|(>=))\$((<)|(<=)|(>)|(>=))[0-9]+$`)
	spilt, _ := regexp.Compile(`(<=)|(>=)|(<)|(>)`)
	if !e.MatchString(expr) {
		return f
	}
	ops := spilt.FindAllString(expr, -1)
	numbers := spilt.Split(expr, -1)
	if (ops[0] == "<" || ops[0] == "<=") && (ops[1] == "<" || ops[1] == "<=") {
		pre, _ := strconv.Atoi(numbers[0])
		sub, _ := strconv.Atoi(numbers[2])
		if pre <= sub {
			f.GreatThan(pre)
			f.LessThan(sub)
		}
	}

	if (ops[0] == ">" || ops[0] == ">=") && (ops[1] == ">" || ops[1] == ">=") {
		pre, _ := strconv.Atoi(numbers[0])
		sub, _ := strconv.Atoi(numbers[2])
		if pre >= sub {
			f.GreatThan(sub)
			f.LessThan(pre)
		}
	}
	return f
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
