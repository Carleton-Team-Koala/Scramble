package controllers

import (
	"encoding/json"
	"net/http"

	"Scramble/backend/pkg/models"

	"github.com/gorilla/mux"
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Scramble!",
	}
	json.NewEncoder(w).Encode(response)
}

// API endpoint to create new game
// TODO: Figure out how to return this information for frontend
func CreateGame(w http.ResponseWriter, r *http.Request) {
	// TODO: Connect with FrontEnd
	playerName := "player1"

	newGame := models.CreateGame(playerName)
	// res,_ := json.Marshal(newGame)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
	json.NewEncoder(w).Encode(newGame)
}

// API endpoint to join game using unique ID
// TODO: Figure out how to return this information for frontend
func JoinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	// TODO: Connect with FrontEnd
	playerName := "anotherPlayer"
	gameDetails := models.JoinGame(gameID, playerName)
	// res, _ := json.Marshal(gameDetails)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
	json.NewEncoder(w).Encode(gameDetails)
}
