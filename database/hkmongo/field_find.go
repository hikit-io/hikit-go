package hkmongo

import (
	"regexp"
	"strconv"
	"strings"
)

func (f *Field) LessThan(val Any) *Field {
	return f.op(FindOp.LessThan, OpTypeFind, val)
}

func (f *Field) LessThanEqual(val Any) *Field {
	return f.op(FindOp.LessThanEqual, OpTypeFind, val)
}

func (f *Field) GreatThan(val Any) *Field {
	return f.op(FindOp.GreatThan, OpTypeFind, val)
}

func (f *Field) GreatThanEqual(val Any) *Field {
	return f.op(FindOp.GreatThanEqual, OpTypeFind, val)
}

func (f *Field) NotEqual(val Any) *Field {
	return f.op(FindOp.NotEqual, OpTypeFind, val)
}

func (f *Field) All(val Any) *Field {
	return f.op(FindOp.All, OpTypeFind, val)
}

// In (hktypes.IsSlice(val) == true)
func (f *Field) In(val Any) *Field {
	return f.op(FindOp.In, OpTypeFind, val)
}

// NotIn (hktypes.IsSlice(val) == true)
func (f *Field) NotIn(val Any) *Field {
	return f.op(FindOp.NotIn, OpTypeFind, val)
}

func (f *Field) Size(val Ui) *Field {
	return f.op(FindOp.Size, OpTypeFind, val)
}

func (f *Field) Exists(val Any) *Field {
	return f.op(FindOp.Exists, OpTypeFind, val)
}

func (f *Field) Type(val Any) *Field {
	return f.op(FindOp.Type, OpTypeFind, val)
}

func (f *Field) Mod(val Any) *Field {
	return f.op(FindOp.Mod, OpTypeFind, val)
}

func (f *Field) Regex(val Str) *Field {
	return f.op(FindOp.Regex, OpTypeFind, val)
}

func (f *Field) Text(val Any) *Field {
	return f.op(FindOp.Text, OpTypeFind, val)
}

func (f *Field) ElemMatch(val MustStructOrPtr) *Field {
	return f.op(FindOp.ElemMatch, OpTypeFind, val)
}

func (f *Field) Equal(val Any) *Field {
	return f.op(FindOp.Equal, OpTypeFind, val)
}

func (f *Field) Hint(val int) *Field {
	return f.op(QueryOp.Hint, OpTypeQuery, val)
}

func (f *Field) Max(val Any) *Field {
	return f.op(UpdateOp.Max, OpTypeUpdate, val)
}

func (f *Field) Expression(expr string) *Field {
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

func (f *Field) Sort(val SortType) *Field {
	return f.op(QueryOp.Sort, OpTypeQuery, val)
}

func (f *Field) Projection(val bool) *Field {
	return f.op(QueryOp.Projection, OpTypeQuery, val)
}
