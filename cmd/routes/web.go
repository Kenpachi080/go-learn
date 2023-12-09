package routes

import (
	"github.com/gorilla/mux"
	"project-go/cmd/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", controllers.Index)

	return route
}
