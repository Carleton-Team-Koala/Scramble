package models

import (
	"github.com/dchest/uniuri"
)

var GameList map[string]Game

type Game struct {
	GameID string `json:"GameID"`
	Board [15][15]string `json:"Board"`
	Players []Player `json:"Players"`
}

type Player struct {
	Name string `json:"name"`
	Score int `json:"score"`
}

func generateNewGameID() string{
	gameID := uniuri.NewLen(6)
	return gameID
}

func CreateGame(playerName string) *Game{   
    gameID := generateNewGameID()
	newPlayer := Player{
		Name: playerName,
		Score: 0,
	}
	playerList := []Player{newPlayer}
	newGame := Game{
		GameID: gameID,
		Board: [15][15]string{},
		Players: playerList,
	}
	
	if GameList == nil {
		GameList = make(map[string]Game)
	}

	GameList[gameID] = newGame
	
	return &newGame
}

func JoinGame(gameID string, playerName string) *Game {
	loadGame := GetGameById(gameID)
	newPlayer := Player{
		Name: playerName,
		Score: 0,
	}
	loadGame.Players = append(loadGame.Players, newPlayer)

	return &loadGame
}

func GetGameById(gameID string) Game{
	return GameList[gameID]
}
