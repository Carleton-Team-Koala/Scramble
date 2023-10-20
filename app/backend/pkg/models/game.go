package models

import (
	"fmt"

	"github.com/dchest/uniuri"
)

func generateNewGameID() string {
	gameID := uniuri.NewLen(6)
	return gameID
}

func GetRandomTile(gameID string) string {
	// get game from GameList
	loadGame := GetGameById(gameID)
	var keys []string
	// get list of tiles that are available
	for k := range loadGame.AvailableLetters {
		if loadGame.AvailableLetters[k] > 0 {
			keys = append(keys, k)
		}
	}
	// in golang, iteration order is not specified and is not guaranteed to be the same from one iteration to the next
	// this will therefore return a random value
	return keys[0]
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

	newLetterDistribution := getNewLetterDistribution()

	// create new game struct with all the new information
	newGame := Game{
		GameID:           gameID,
		Board:            [15][15]string{},
		AvailableLetters: newLetterDistribution,
		Players:          playerList,
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
	loadGame := GetGameById(gameID)

	// create new player
	newPlayer := Player{
		Name:  playerName,
		Score: 0,
	}

	// add new player to player list
	loadGame.Players = append(loadGame.Players, newPlayer)

	return loadGame
}

// Load Game by GameID
func GetGameById(gameID string) *Game {
	if checkGameExists(gameID) != nil {
		return nil
	}
	loadedGame := GameList[gameID]

	return &loadedGame
}

// Update the Board with player's move
func UpdateBoard(gameID string, playerMove Move) {
	loadedGame := GetGameById(gameID)

	if !ValidateMove(playerMove, gameID) {
		// TODO: Change response to something else that makes more sense
		fmt.Println("Invalid Move")
	}

	// update board state
	loadedGame.Board[playerMove.XLoc][playerMove.YLoc] = playerMove.Letter

	// update available tiles status
	loadedGame.AvailableLetters[playerMove.Letter] -= 1

	// TODO: Only used to replace Game in GameList. Remove this once database is connected
	GameList[gameID] = *loadedGame

	// TODO: Only used for debugging purposes. Remove this later
	fmt.Println(loadedGame)
}
