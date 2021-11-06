package hkmongo

import (
	"testing"
)

type User struct {
}

func TestDatabase_Col(t *testing.T) {
	cli := Database{
		dbname: "",
		Client: nil,
		tables: map[TableName]*Collection{},
	}
	cli.Col(User{})
}
