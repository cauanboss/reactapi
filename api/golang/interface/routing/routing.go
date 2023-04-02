package routing

import (
	"api/core/user"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	app := mux.NewRouter()

	userController := app.PathPrefix("/user").Subrouter()
	user.UserController(userController)
	return app
}
