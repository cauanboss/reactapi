package model

import "gopkg.in/mgo.v2/bson"

// import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"_id" bson: "_id"`
	Name     string        `json:"name" bson: "name"`
	Gender   string        `json:"gender" bson: "gender"`
	Age      int32         `json:"age" bson: "age"`
	Username string        `json:"username" bson: "username"`
	Password string        `json:"password" bson: "password"`
}
