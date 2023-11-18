package routes

import (
	"github.com/gorilla/mux"

	"Scramble/app/backend/pkg/controllers"
)

// list of routes for application
var RegisterRoutes = func(router *mux.Router, appInterface controllers.AppControllerInterface) {
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/newgame/", appInterface.AppCreateGame).Methods("POST")
	router.HandleFunc("/joingame/{gameID}/", appInterface.AppJoinGame).Methods("POST")
	router.HandleFunc("/{gameID}/updategame/", appInterface.AppUpdateMove).Methods("POST")
	router.HandleFunc("/startgame/{gameID}/", appInterface.AppStartGame).Methods("GET")
	router.HandleFunc("/refreshhand/{gameID}/", appInterface.AppRefreshHand).Methods("POST")
	router.HandleFunc("/getgamestate/{gameID}/", appInterface.AppReturnGameState).Methods("GET")
	router.HandleFunc("/skipturn/{gameID}/", appInterface.AppSkipTurn).Methods("POST")
	router.HandleFunc("/resigngame/{gameID}/", appInterface.AppResignGame).Methods("POST")
}
