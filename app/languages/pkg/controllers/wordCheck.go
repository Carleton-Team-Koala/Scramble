// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package controllers

import (
	"encoding/json"
	"net/http"

	"Scramble/app/languages/pkg/models"

	"github.com/gorilla/mux"
)

// HomePage is a handler function for the home page of the Languages API.
// It writes a JSON response with a welcome message to the http.ResponseWriter.
func HomePage(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Languages API!",
	}
	json.NewEncoder(w).Encode(response)
}

// LetterScores is a handler function that calculates the score of a given letter.
// It takes an HTTP response writer and request as parameters.
// The letter to be scored is extracted from the request URL path.
// The function retrieves the score of the letter using the GetLetterScore function from the models package.
// The score is then encoded as JSON and sent as the response.
func LetterScores(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inputLetter := vars["letter"]

	letterScore := models.GetLetterScore(inputLetter)

	json.NewEncoder(w).Encode(letterScore)
}

// getter
// LetterDistribution is a handler function that returns the distribution of letters in the alphabet as a JSON response.
// It marshals the alphabet distribution data into JSON format and writes it to the response writer.
// If marshaling fails, it returns an HTTP 500 Internal Server Error.
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

// WordCheck checks if a given word is valid based on the letters it contains.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The word to be checked is extracted from the request's URL parameters.
// It calls the CheckLetter function from the models package to validate the word.
// The result is encoded as JSON and sent back in the http.ResponseWriter.
func WordCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	isValidWord := models.CheckLetter(word)

	json.NewEncoder(w).Encode(isValidWord)
}
