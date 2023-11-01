package models

import (
	"errors"
	"fmt"

	"github.com/dchest/uniuri"
)

type App struct {
	LanguageClient LanguageClient
	DatabaseClient DatabaseClient
}

type AppInterface interface {
	CreateGame(playerName string) (string, error)
	JoinGame(gameID string, playerName string) error
	StartGame(gameID string) (*Game, error)
	GetGameById(gameID string) (*Game, error)
	UpdateGameState(gameID string, playerMove []Move, playerName string) (*Game, error)
	ValidateMove(playerMove Move, playerName string, gameID string) bool
}

// create new game struct
func (app *App) CreateGame(playerName string) (string, error) {
	gameID := ""

	// generate new game id until unique ID is made
	for {
		gameID = generateNewGameID()
		exists, _ := app.DatabaseClient.CheckGameExists(gameID)
		if !(*exists) {
			break
		}
	}

	// Create new player with input name
	newPlayer := PlayerInfo{
		Score: 0,
		Hand: []string{},
	}

	// add player to player list
	playerList := make(map[string]PlayerInfo)
	playerList[playerName] = newPlayer

	newLetterDistribution := app.LanguageClient.GetNewLetterDistribution()

	// create new game struct with all the new information
	newGame := Game{
		GameID:           gameID,
		Board:            [15][15]string{},
		AvailableLetters: newLetterDistribution,
		Players:          playerList,
	}


	// Add game to database
	err := app.DatabaseClient.AddNewGameToDB(newGame)
	if err != nil {
		return "", err
	}

	return newGame.GameID, nil
}

// add player to already existing game
func (app *App) JoinGame(gameID string, playerName string) error {
	// get game from GameList
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		return err
	}

	// create new player
	newPlayer := PlayerInfo{
		Score: 0,
		Hand: []string{},
	}

	// add new player to player list
	loadGame.Players[playerName] = newPlayer

	app.DatabaseClient.UpdateGameToDB(gameID, *loadGame)

	return nil
}

// start game
func (app *App) StartGame(gameID string) (*Game, error) {
	// get game from GameList
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		return nil, err
	}

	for player, _ := range loadGame.Players {
		var randomStartingTiles []string
		for i := 0; i < 7; i++ {     
			randomTile := getRandomTile(loadGame.AvailableLetters) 
			randomStartingTiles = append(randomStartingTiles, randomTile) 
			loadGame.AvailableLetters[randomTile] -= 1
		}  
		if copyPlayer, ok := loadGame.Players[player]; ok {
			copyPlayer.Hand = randomStartingTiles
			loadGame.Players[player] = copyPlayer
		}

    }
	app.DatabaseClient.UpdateGameToDB(gameID, *loadGame)

	return loadGame, nil
}

// Load Game by GameID
func (app *App) GetGameById(gameID string) (*Game, error) {
	exists, _ :=  app.DatabaseClient.CheckGameExists(gameID)
	if !(*exists) {
		return nil, nil
	}
	loadedGame, err := app.DatabaseClient.GetGameByGameID(gameID)
	if err != nil{
		fmt.Println(fmt.Errorf("%w", err))
		return nil, err
	}

	return loadedGame, nil
}

func (app *App) UpdateGameState(gameID string, playerMove []Move, playerName string) (*Game, error) {
	loadedGame, err := app.GetGameById(gameID)
	if err != nil{
		return nil, err
	}
	
	var randomTiles []string
	var randomTile string

	// update the board once every move is validated and get random tiles to replace tiles used
	for _, move := range playerMove {
		if app.ValidateMove(move, playerName, gameID) {
			loadedGame, randomTile = updateBoardAndHand(*loadedGame, move, playerName)
			randomTiles = append(randomTiles, randomTile)
			} else {
			return nil, errors.New("Invalid Move")
		}
		
	}

	// TODO: get score for entered word
	wordScore := 10
	// wordScore, err := app.LanguageClient.scoring(*loadedGame, playerMove)
	// if err != nil {
	// 	return nil, err
	// }

	loadedGame.Players, err = updateScore(wordScore, loadedGame.Players, playerName)

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	return loadedGame, nil
}

func generateNewGameID() string {
	gameID := uniuri.NewLen(6)
	return gameID
}

func getRandomTile(availableLetters map[string]int) string {
	var keys []string
	// get list of tiles that are available
	for k := range availableLetters {
		if availableLetters[k] > 0 {
			keys = append(keys, k)
		}
	}
	// in golang, iteration order is not specified and is not guaranteed to be the same from one iteration to the next
	// this will therefore return a random value
	return keys[0]
}

// Update the Board with player's move
func updateBoardAndHand(loadedGame Game, playerMove Move, playerName string) (*Game, string) {
	// update board state
	loadedGame.Board[playerMove.XLoc][playerMove.YLoc] = playerMove.Letter

	randomTile := getRandomTile(loadedGame.AvailableLetters)
	index := 0
	for _, iterateLetter := range loadedGame.Players[playerName].Hand {
		if iterateLetter == playerMove.Letter {
			loadedGame.Players[playerName].Hand[index] = randomTile
			break
		}
		index++
	}

	return &loadedGame, randomTile
}

// Update Player's Scores
func updateScore(wordScore int, currPlayers map[string]PlayerInfo, currPlayer string) (map[string]PlayerInfo, error) {
	// calculate new score
	currScore := currPlayers[currPlayer].Score
	newScore := currScore + wordScore

	// update player score to struct
	if copyPlayer, ok := currPlayers[currPlayer]; ok {
		copyPlayer.Score = newScore
		currPlayers[currPlayer] = copyPlayer
	}

	return currPlayers, nil
}
