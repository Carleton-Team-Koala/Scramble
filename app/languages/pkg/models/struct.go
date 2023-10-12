package models

// Alphabet datatype - stores an alphabet and its associated
type Alphabet struct {
	ListOfWords map[string]int `json:"listOfWords"`
}

type Dictionary struct {
	WordList []string
}
