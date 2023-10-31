package controllers

import (
	"encoding/json"
	"errors"
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
	AppStartGame(w http.ResponseWriter, r *http.Request)
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
	// unmarshal json response
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var createGameResp models.PlayerNameResp
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&createGameResp)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}


	newGameID, err := a.AppInterface.CreateGame(createGameResp.PlayerName)
	if err != nil {
		errorResponse(w, "Not able to create new game", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newGameID)
}

// API endpoint to join game using unique ID
// TODO: Figure out how to return this information for frontend
func (a *AppController) AppJoinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	
	// unmarshal json response
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var joinGameResp models.PlayerNameResp
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&joinGameResp)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	
	err = a.AppInterface.JoinGame(gameID, joinGameResp.PlayerName)
	if err != nil {
		errorResponse(w, "Not able to join game", http.StatusBadRequest)
		return
	}
}

// TODO: Figure out how to return this information for frontend
func (a *AppController) AppStartGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	gameDetails := a.AppInterface.StartGame(gameID)

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

	var listOfMoves models.UpdateGameResp
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

	var randomTiles []string

	// update the board once every move is validated and get random tiles to replace tiles used
	for _, move := range listOfMoves.Updates {
		randomTile := a.AppInterface.UpdateBoard(gameID, move, listOfMoves.PlayerName)
		randomTiles = append(randomTiles, *randomTile)
	}

	// TODO: Update Game Score

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
