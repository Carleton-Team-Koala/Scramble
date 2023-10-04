package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Scramble/backend/pkg/models"

	"github.com/gorilla/mux"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var playerName string
	fmt.Scanln(&playerName)

	newGame := models.CreateGame(playerName)
	res,_ := json.Marshal(newGame)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	gameDetails := models.GetGameById(gameID)
	res, _ := json.Marshal(gameDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}