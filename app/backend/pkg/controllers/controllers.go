package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"Scramble/app/backend/pkg/models"

	"github.com/gorilla/mux"
)

type AppController struct {
	AppInterface models.App
}

type AppControllerInterface interface {
	AppCreateGame(w http.ResponseWriter, r *http.Request)
	AppJoinGame(w http.ResponseWriter, r *http.Request)
	AppUpdateMove(w http.ResponseWriter, r *http.Request)
}

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Scramble!",
	}
	json.NewEncoder(w).Encode(response)
}

// API endpoint to create new game
// TODO: Figure out how to return this information for frontend
func (a *AppController) AppCreateGame(w http.ResponseWriter, r *http.Request) {
	// TODO: Connect with FrontEnd
	playerName := "player1"

	newGame := a.AppInterface.CreateGame(playerName)

	json.NewEncoder(w).Encode(newGame)
}

// API endpoint to join game using unique ID
// TODO: Figure out how to return this information for frontend
func (a *AppController) AppJoinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	// TODO: Connect with FrontEnd
	playerName := "anotherPlayer"
	gameDetails := models.JoinGame(gameID, playerName)

	json.NewEncoder(w).Encode(gameDetails)
}

func (a *AppController) AppUpdateMove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]

	// unmarshal json response
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var listOfMoves models.Resp
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&listOfMoves)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	// validate each move
	for _, move := range listOfMoves.Updates {
		if !models.ValidateMove(move, gameID) {
			fmt.Println("invalid move")
			return
		}
	}

	// update the board once every move is validated
	for _, move := range listOfMoves.Updates {
		models.UpdateBoard(gameID, move)
	}

	// TODO: Update Game Score

	// create random tiles to replace the used tiles
	// this should be done after board gets updated so that we use the remaining tiles only
	var randomTiles []string
	for i := 0; i < len(listOfMoves.Updates); i++ {
		randomTile := models.GetRandomTile(gameID)
		randomTiles = append(randomTiles, randomTile)
	}

	json.NewEncoder(w).Encode(randomTiles)
}

// Error response
func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
