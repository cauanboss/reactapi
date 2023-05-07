package model

// import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Name     string `json:"name" bson: "name"`
	Gender   string `json:"gender" bson: "gender"`
	Age      int32  `json:"age" bson: "age"`
	Username string `json:"username" bson: "username"`
	Password string `json:"password" bson: "password"`
}
