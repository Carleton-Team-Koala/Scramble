package models

var GameList map[string]Game


type Game struct {
	GameID  string         `json:"GameID"`
	Board   [15][15]string `json:"Board"`
	Players []Player       `json:"Players"`
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}