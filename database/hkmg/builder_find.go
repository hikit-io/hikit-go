package hkmg

var FindOp = findOp{
	Equal:          "$eq",
	GreatThan:      "$gt",
	GreatThanEqual: "$gte",
	In:             "$in",
	LessThan:       "$lt",
	LessThanEqual:  "$lte",
	NotEqual:       "$ne",
	NotIn:          "$nin",
	All:            "$all",
	ElemMatch:      "$elemMatch",
	Size:           "$size",
	Exists:         "$exists",
	Type:           "$type",
	Mod:            "$mod",
	Regex:          "$regex",
	Text:           "$Text",
}

type findOp struct {
	LessThan       string
	LessThanEqual  string
	GreatThan      string
	GreatThanEqual string
	NotEqual       string
	All            string
	In             string
	NotIn          string
	Size           string
	Exists         string
	Type           string
	Mod            string
	Regex          string
	Text           string
	ElemMatch      string
	Equal          string
}
