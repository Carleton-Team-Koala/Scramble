// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package controllers

import (
	"encoding/json"
	"net/http"

	"Scramble/app/languages/pkg/models"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Languages API!",
	}
	json.NewEncoder(w).Encode(response)
}

// getter
func LetterScores(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inputLetter := vars["letter"]

	letterScore := models.GetLetterScore(inputLetter)

	json.NewEncoder(w).Encode(letterScore)
}

// getter
func LetterDistribution(w http.ResponseWriter, r *http.Request) {
	distribution := models.AlphabetDistribution

	jsonData, err := json.Marshal(distribution)
	if err != nil {
		// Handle the error if marshaling fails
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") //approves response
	w.WriteHeader(http.StatusOK)                       //good HTTP response
	w.Write(jsonData)
}

// check if input word is valid
func WordCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	isValidWord := models.CheckLetter(word)

	json.NewEncoder(w).Encode(isValidWord)
}
