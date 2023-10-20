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
