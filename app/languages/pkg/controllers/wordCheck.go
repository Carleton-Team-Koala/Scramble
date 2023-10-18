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

var dictionaryText = "app/languages/pkg/controllers/englishWordList.txt"

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

	// importJSONdata(alphabetDistributionFilePath, distribution)
	// alphabetReturner(w, r, distribution)
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
