package routes

import (
	"github.com/gorilla/mux"

	"Scramble/app/backend/pkg/controllers"
)

// list of routes for application
var RegisterRoutes = func(router *mux.Router, appInterface controllers.AppControllerInterface) {
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/newgame/", appInterface.AppCreateGame).Methods("GET")
	router.HandleFunc("/joingame/{gameID}/", appInterface.AppJoinGame).Methods("GET")
	router.HandleFunc("/{gameID}/updategame/", appInterface.AppUpdateMove).Methods("POST")
	router.HandleFunc("/startgame/{gameID}/", appInterface.AppStartGame).Methods("GET")
}
