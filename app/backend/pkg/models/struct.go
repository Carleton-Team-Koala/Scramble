package models

var GameList map[string]Game

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

type Move struct {
	Letter string `json:"letter"`
	XLoc   int    `json:"xLoc"`
	YLoc   int    `json:"yLoc"`
}

type Resp struct {
	PlayerName string `json:"playerName"`
	Updates    []Move `json:"updates"`
}
