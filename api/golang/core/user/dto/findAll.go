package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type FindUser struct {
	Id     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name   string             `json:"name" bson: "name"`
	Gender string             `json:"gender" bson: "gender"`
	Age    int32              `json:"age" bson: "age"`
}
