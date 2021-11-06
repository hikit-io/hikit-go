package hkmongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	DeleteTime *time.Time         `bson:"delete_time" json:"delete_time"`
	CreateTime *time.Time         `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time         `bson:"update_time" json:"update_time"`
}
