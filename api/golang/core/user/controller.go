package user

import (
	"api/core/user/dto"
	"api/core/user/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type MuxApp struct {
	app *mux.Router
}

func UserController(app *mux.Router) *MuxApp {
	app.HandleFunc("/{name}", findOne).Methods("GET")
	app.HandleFunc("", findAll).Methods("GET")
	app.HandleFunc("", createUser).Methods("POST")
	app.HandleFunc("/{id}", deleteOne).Methods("DELETE")
	return &MuxApp{app}
}

func findOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	name := mux.Vars(r)["name"]
	filter := bson.M{"name": name}
	u := UserService().FindOne(filter)
	json.NewEncoder(w).Encode(u)
}

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Print("asd")
	u := UserService().FindAll()
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CreateUser")
	w.Header().Set("content-type", "application/json")
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), 500)
		return
	}
	UserService().CreateUser(u)
	json.NewEncoder(w).Encode(u)
}

func deleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	filter := dto.Delete{Id: id}
	result, err := UserService().DeleteOne(filter)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(result)
}
