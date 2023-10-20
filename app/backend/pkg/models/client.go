package models

import (
	"encoding/json"
	"io"
	"net/http"
)

// connect to languages to get the list of available tiles when creating game
func getNewLetterDistribution() map[string]int {
	var letterDistribution map[string]int

	req, err := http.NewRequest("GET", "http://languages:8000/letterDistribution", nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &letterDistribution)

	return letterDistribution
}

// connect to languages to check whether input word is valid
func checkValidWord(inputWord string) bool {
	var isValidWord bool

	req, err := http.NewRequest("GET", "http://languages:8000/checkWord/" + string(inputWord), nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &isValidWord)

	return isValidWord
}