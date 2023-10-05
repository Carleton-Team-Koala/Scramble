package models

import (
	"errors"

	"github.com/dchest/uniuri"
)

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

func generateNewGameID() string {
	gameID := uniuri.NewLen(6)
	return gameID
}

// create new game struct
func CreateGame(playerName string) *Game {
	gameID := ""

	// generate new game id until unique ID is made
	for {
		gameID = generateNewGameID()
		_, exists := GameList[gameID]
		if !exists {
			break
		}
	}

	// Create new player with input name
	newPlayer := Player{
		Name:  playerName,
		Score: 0,
	}

	// add player to player list
	playerList := []Player{newPlayer}

	// create new game struct with all the new information
	newGame := Game{
		GameID:  gameID,
		Board:   [15][15]string{},
		Players: playerList,
	}

	// if GameList does not exist, make a new map
	if GameList == nil {
		GameList = make(map[string]Game)
	}

	// add created game to GameList
	GameList[gameID] = newGame

	return &newGame
}

// add player to already existing game
func JoinGame(gameID string, playerName string) *Game {
	// get game from GameList
	loadGame, _ := GetGameById(gameID)

	// create new player
	newPlayer := Player{
		Name:  playerName,
		Score: 0,
	}

	// add new player to player list
	loadGame.Players = append(loadGame.Players, newPlayer)

	return loadGame
}

func GetGameById(gameID string) (*Game, error) {
	_, exists := GameList[gameID]
	if !exists {
		return nil, errors.New("Game ID not found")
	}

	loadedGame := GameList[gameID]

	return &loadedGame, nil
}
