package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dictionaryText = "app/languages/pkg/controllers/englishWordList.txt"

func GetLetterScore(letter string) int {
	return alphabetScores[letter]
}

// importDict is a function that imports a dictionary from a JSON file located at the given textPath.
// It returns a pointer to a Dictionary struct containing the imported word list.
// If the file cannot be opened, it returns nil.
func importDict(textPath string) *Dictionary {
	var wordList = new(Dictionary)

	// Open the JSON file
	file, err := os.Open(textPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
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

	return wordList
}

// CheckLetter is a function that checks if a given word is valid based on the letters it contains.
// It takes in a string as a parameter and returns a boolean value.
// It imports a dictionary using the importDict function and then performs a binary search on the dictionary to check if the word is valid.
// If the word is valid, it returns true. Otherwise, it returns false.
func CheckLetter(searchWord string) bool {
	wordList := importDict(dictionaryText)

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
