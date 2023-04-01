package user

import (
	"api/core/user/dto"
	"api/core/user/model"
	"api/interface/database/mongodb"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MuxApp struct {
	app *mux.Router
}

type UserClient struct {
	client *mongo.Collection
}

func NewUserClient(ur *mongo.Collection) *UserClient {
	return &UserClient{ur}
}

func App(app *mux.Router) *MuxApp {

	nu := NewUserClient(mongodb.GetClient())

	app.HandleFunc("/{name}", nu.findOne).Methods("GET")
	app.HandleFunc("", nu.findAll).Methods("GET")
	app.HandleFunc("", nu.createUser).Methods("POST")
	app.HandleFunc("/{id}", nu.deleteOne).Methods("DELETE")
	return &MuxApp{app}
}

func (ur *UserClient) findOne(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	name := mux.Vars(r)["name"]
	filter := bson.M{"name": name}
	u := model.User{}
	ur.client.FindOne(ctx, filter).Decode(&u)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (ur *UserClient) findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	ctx := context.TODO()
	users := []dto.FindUser{}
	cursor, err := ur.client.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)
}

func (ur *UserClient) createUser(w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)

	fmt.Println("No Err", err, u)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), 500)
		return
	}
	ur.client.InsertOne(context.TODO(), u)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (ur *UserClient) deleteOne(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	id := mux.Vars(r)["id"]
	filter := bson.M{"_id": id}
	result, err := ur.client.DeleteOne(ctx, filter)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(result)
}
