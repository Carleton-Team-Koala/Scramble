// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"Scramble/app/languages/pkg/models"

	"github.com/gorilla/mux"
)

var wordList = new(models.Dictionary)

var alphabetScoresFilePath = "../../pkg/controllers/englishAlphabetScores.json"
var alphabetDistributionFilePath = "../../pkg/controllers/englishAlphabetDistribution.json"
var dictionaryText = "../../pkg/controllers/englishWordList.txt"

// getter
func LetterScores(w http.ResponseWriter, r *http.Request) {
	var scores = new(models.Alphabet)
	importJSONdata(alphabetScoresFilePath, scores)
	alphabetReturner(w, r, scores)
}

// getter
func LetterDistribution(w http.ResponseWriter, r *http.Request) {
	var distribution = new(models.Alphabet)
	importJSONdata(alphabetDistributionFilePath, distribution)
	alphabetReturner(w, r, distribution)
}

// returns a JSON file of a given alphabet struct
func alphabetReturner(w http.ResponseWriter, r *http.Request, activeAlphabet *models.Alphabet) {
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

func importJSONdata(path string, activeAlphabet *models.Alphabet) {
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

func importDict(textPath string, words *models.Dictionary) {
	// Open the JSON file
	file, err := os.Open(textPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		wordList.WordList = append(wordList.WordList, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
	}
}

func checkLetter(searchWord string) bool {
	importDict(dictionaryText, wordList)

	searchWord = strings.ToLower(searchWord)
	left, right := 0, len(wordList.WordList)-1

	for left <= right {
		mid := left + (right-left)/2
		midWord := strings.ToLower(wordList.WordList[mid])

		if midWord == searchWord {
			return true
		} else if midWord < searchWord {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func WordCheck(w http.ResponseWriter, r *http.Request) {
	jsonObject := make(map[string]interface{})

	w.Header().Set("Content-Type", "application/json") //approves response
	w.WriteHeader(http.StatusOK)                       //good HTTP response
	word := mux.Vars(r)["word"]

	if checkLetter(word) {
		jsonObject["result"] = "true"
	} else {
		jsonObject["result"] = "false"
	}
	jsonData, err := json.Marshal(jsonObject)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w.Write(jsonData)
}
