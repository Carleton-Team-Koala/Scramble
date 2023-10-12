package routes

import (
	"github.com/gorilla/mux"

	"Scramble/app/backend/pkg/controllers"
)

// list of routes for application
var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/newgame/", controllers.CreateGame).Methods("GET")
	router.HandleFunc("/joingame/{gameID}/", controllers.JoinGame).Methods("GET")
	router.HandleFunc("/{gameID}/updategame/", controllers.UpdateMove).Methods("POST")
}