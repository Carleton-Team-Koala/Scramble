package routes

import (
	"github.com/gorilla/mux"

	"Scramble/backend/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/newGame/", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/joinGame/{gameID}", controllers.JoinGame).Methods("GET")
}