package models

import (
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
	UpdateBoard(gameID string, playerMove Move)
	GetRandomTile(gameID string) string
	GetGameById(gameID string) *Game
	ValidateMove(playerMove Move, gameID string) bool
}

func generateNewGameID() string {
	gameID := uniuri.NewLen(6)
	return gameID
}

func (app *App) GetRandomTile(gameID string) string {
	// get game from GameList
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	}
	var keys []string
	// get list of tiles that are available
	for k := range loadGame.AvailableLetters {
		if loadGame.AvailableLetters[k] > 0 {
			keys = append(keys, k)
			loadGame.AvailableLetters[k] -= 1
		}
	}
	// in golang, iteration order is not specified and is not guaranteed to be the same from one iteration to the next
	// this will therefore return a random value
	return keys[0]
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
func (app *App) StartGame(gameID string) *Game {
	var randomStartingTiles []string

	// get game from GameList
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	}

	for player, _ := range loadGame.Players {
		for i := 0; i < 7; i++ {              
			randomTile := app.GetRandomTile(gameID) 
			randomStartingTiles = append(randomStartingTiles, randomTile)  
		}  
		if copyPlayer, ok := loadGame.Players[player]; ok {
			copyPlayer.Hand = randomStartingTiles
			loadGame.Players[player] = copyPlayer
		}

    }
	app.DatabaseClient.UpdateGameToDB(gameID, *loadGame)

	return loadGame
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

// Update the Board with player's move
func (app *App) UpdateBoard(gameID string, playerMove Move, playerName string) *string {
	loadedGame, err := app.GetGameById(gameID)
	if err != nil{
		fmt.Println(fmt.Errorf("%w", err))
		return nil
	}

	if !app.ValidateMove(playerMove, playerName, gameID) {
		// TODO: Change response to something else that makes more sense
		fmt.Println("Invalid Move")
		return nil
	}

	// update board state
	loadedGame.Board[playerMove.XLoc][playerMove.YLoc] = playerMove.Letter

	// update available tiles status
	loadedGame.AvailableLetters[playerMove.Letter] -= 1

	randomTile := app.GetRandomTile(gameID)
	index := 0
	for _, iterateLetter := range loadedGame.Players[playerName].Hand {
		if iterateLetter == playerMove.Letter {
			loadedGame.Players[playerName].Hand[index] = iterateLetter
			break
		}
		index++
	}

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	// TODO: Only used for debugging purposes. Remove this later
	fmt.Println(loadedGame)

	return &randomTile
}