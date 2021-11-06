package hkmongo

var QueryOp = queryOp{
	//Limit:      "limit",
	Comment:    "comment",
	Projection: "projection",
	Sort:       "sort",
	Hint:       "hint",
}

type queryOp struct {
	//Limit      string
	//Skip       string
	Sort    string
	Comment string
	//Max        string
	//Min        string
	//Hint       string
	Hint       string
	Projection string
}
