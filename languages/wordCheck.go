// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var scores = new(alphabet)
var distribution = new(alphabet)
var wordList = new(dictionary)

// Alphabet datatype - stores an alphabet and its associated
type alphabet struct {
	ListOfWords map[string]int `json:"listOfWords"`
}

type dictionary struct {
	wordList []string
}

// getter
func letterScores(w http.ResponseWriter, r *http.Request) {
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

func importDict(textPath string, words *dictionary) {
	file, err := os.Open(textPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		wordList.wordList = append(wordList.wordList, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
	}
}

func checkLetter(searchWord string) bool {
	searchWord = strings.ToLower(searchWord)
	left, right := 0, len(wordList.wordList)-1

	for left <= right {
		mid := left + (right-left)/2
		midWord := strings.ToLower(wordList.wordList[mid])

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

func wordCheckInterface(w http.ResponseWriter, r *http.Request) {
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

func main() {

	var alphabetScoresFilePath = "englishAlphabetScores.json"
	var alphabetDistributionFilePath = "englishAlphabetDistribution.json"
	var dictionaryText = "englishWordList.txt"

	importJSONdata(alphabetScoresFilePath, scores)
	importJSONdata(alphabetDistributionFilePath, distribution)
	importDict(dictionaryText, wordList)

	r := mux.NewRouter()
	r.HandleFunc("/checkWord/{word}", wordCheckInterface).Methods("GET")

	// TODO: Import data and other setup code here

	// Define additional routes for your application
	r.HandleFunc("/letterScores", letterScores)
	r.HandleFunc("/letterDistribution", letterDistribution)

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
