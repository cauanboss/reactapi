package routing

import (
	"api/core/user"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	app := mux.NewRouter()

	userApp := app.PathPrefix("/user").Subrouter()
	user.App(userApp)
	return app
}
