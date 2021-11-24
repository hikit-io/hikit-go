package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestAdd(t *testing.T) {
	users := []interface{}{
		User{
			ID:   primitive.NewObjectID(),
			Name: "nieaowei",
			Age:  10,
			Addr: "hunan",
		},
		User{
			ID:   primitive.NewObjectID(),
			Name: "nieaowei1",
			Age:  20,
			Addr: "ew",
		},
		User{
			ID:   primitive.NewObjectID(),
			Name: "nieaowei3",
			Age:  32,
			Addr: "ew",
		},
	}
	db.Col(User{}).HInsertMany(ctx, users)
}
