package models


type Game struct {
	GameID           string         `json:"GameID"`
	Board            [15][15]string `json:"Board"`
	AvailableLetters map[string]int `json:"LetterDistribution"`
	Players          map[string]PlayerInfo       `json:"Players"`
}

type PlayerInfo struct {
	Score int    `json:"score"`
	Hand []string `json:"hand"`
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

type WordScore struct {
	Valid bool
	Score int
}