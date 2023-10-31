package models

import (
	"encoding/json"
	"io"
	"net/http"
)

// connect to languages to get the list of available tiles when creating game
func (c *LanguageClient) GetNewLetterDistribution() map[string]int {
	var letterDistribution map[string]int
	newURL := c.BaseURL + "letterDistribution"

	req, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &letterDistribution)

	return letterDistribution
}

// connect to languages to check whether input word is valid
func (c *LanguageClient) CheckValidWord(inputWord string) bool {
	var isValidWord bool
	newURL := c.BaseURL + "/checkWord/" + inputWord

	req, err := http.NewRequest("GET", newURL, nil)

	if err != nil {
		panic(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &isValidWord)

	return isValidWord
}

// connect to languages to return a letter score
func (c *LanguageClient) GetLetterScore(inputLetter string) int {
	var value int
	newURL := c.BaseURL + "/letterScores/" + inputLetter

	req, err := http.NewRequest("GET", newURL, nil)

	if err != nil {
		panic(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &value)

	return value
}
