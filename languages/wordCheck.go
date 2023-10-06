// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Alphabet datatype - stores an alphabet and its associated
type alphabet struct {
	ListOfWords map[string]int `json:"listOfWords"`
}

// TODO: simplify to make this one alphabet system
// alphabetScoresFilePath and alphabetDistributionFilePath each represent the scores per letter/tile, and how many tiles of each letter/blank characters there are in this alphabet.
var scores = new(alphabet)
var distribution = new(alphabet)
var alphabetScoresFilePath = "englishAlphabetScores.json"
var alphabetDistributionFilePath = "englishAlphabetDistribution.json"

// getter
func letterScores(w http.ResponseWriter, r *http.Request) {
	fmt.Println
	alphabetReturner(w, r, scores)
}

// getter
func letterDistribution(w http.ResponseWriter, r *http.Request) {
	alphabetReturner(w, r, distribution)
}

// returns a JSON file of a given alphabet struct
func alphabetReturner(w http.ResponseWriter, r *http.Request, activeAlphabet *alphabet) {
	// recreate json dataset of active alphabet
	jsonData, err := json.Marshal(activeAlphabet)
	if err != nil {
		// Handle the error if marshaling fails
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") //approves response
	w.WriteHeader(http.StatusOK)                       //good HTTP response
	w.Write(jsonData)
}

func importJSONdata(path string, activeAlphabet *alphabet) {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a decoder to read the JSON data
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the struct
	err = decoder.Decode(&activeAlphabet)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func main() {
	importJSONdata(alphabetScoresFilePath, scores)
	importJSONdata(alphabetDistributionFilePath, distribution)

	//http listener paths
	http.HandleFunc("/letterScores", letterScores)
	http.HandleFunc("/letterDistribution", letterDistribution)

	//run server
	http.ListenAndServe(":8080", nil)
}
