package mongo

var UpdateOp = _UpdateOp{
	Set:              "$set",
	Unset:            "$unset",
	Push:             "$push",
	Pop:              "$pop",
	Pull:             "$pull",
	PullAll:          "$pullAll",
	AddToSet:         "$addToSet",
	PlaceholderAll:   "$[]",
	PlaceholderSome:  "$",
	PlaceholderIndex: "$[%s]",
}

type _UpdateOp struct {
	CurrentDate      string
	Increment        string
	Min              string
	Max              string
	Rename           string
	Set              string
	SetOnInsert      string
	Unset            string //
	PlaceholderAll   string
	PlaceholderSome  string
	PlaceholderIndex string
	AddToSet         string
	Pop              string
	Pull             string
	Push             string
	PullAll          string //
	Each             string
	Position         string
	Slice            string
	Sort             string
}
