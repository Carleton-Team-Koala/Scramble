package routes

import (
	"github.com/gorilla/mux"

	"Scramble/app/languages/pkg/controllers"
)

// list of routes for application
var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/checkWord/{word}", controllers.WordCheck).Methods("GET")
	router.HandleFunc("/letterScores/{letter}", controllers.LetterScores).Methods("GET")
	router.HandleFunc("/letterDistribution", controllers.LetterDistribution)
}
