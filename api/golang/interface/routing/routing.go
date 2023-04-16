package routing

import (
	"api/core/user"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	app := mux.NewRouter()

	app.HandleFunc("/", hello).Methods("GET")
	userController := app.PathPrefix("/user").Subrouter()
	user.UserController(userController)
	return app
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
