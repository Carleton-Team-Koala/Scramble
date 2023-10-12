package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"Scramble/app/backend/pkg/models"

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

func UpdateMove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]

	// unmarshal json response
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var m models.Move
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&m)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Update Game board
	models.UpdateBoard(gameID, m)

	// TODO: Update Game Score

	// TODO: return response(random tile)
	// json.NewEncoder(w).Encode(response)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
