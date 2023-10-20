package models

type Game struct {
	GameID           string         `json:"GameID"`
	Board            [15][15]string `json:"Board"`
	AvailableLetters map[string]int `json:"LetterDistribution"`
	Players          []Player       `json:"Players"`
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
