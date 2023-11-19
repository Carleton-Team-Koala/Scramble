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
		Hand:  []string{},
	}

	// add player to player list
	playerStructList := make(map[string]PlayerInfo)
	playerStructList[playerName] = newPlayer

	playerList := []string{playerName}

	newLetterDistribution := app.LanguageClient.GetNewLetterDistribution()

	// create new game struct with all the new information
	newGame := Game{
		GameID:           gameID,
		Board:            [15][15]string{},
		AvailableLetters: newLetterDistribution,
		Players:          playerStructList,
		CurrentPlayer:    "",
		PlayerList:       playerList,
		TotalMoves:       0,
		GameStarted:      false,
		GameOver:         false,
		Winner:           "",
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

	if playerName == loadGame.PlayerList[0] {
		return errors.New("player already exists in game, try a different name")
	}

	if loadGame.GameStarted {
		return errors.New("cannot join game: game already started")
	}

	if len(loadGame.PlayerList) >= 2 {
		return errors.New("cannot join game: game already has two players")
	}

	if loadGame.GameOver {
		return errors.New("cannot join game: this game is already over")
	}

	// create new player
	newPlayer := PlayerInfo{
		Score: 0,
		Hand:  []string{},
	}

	// add new player to player list
	loadGame.Players[playerName] = newPlayer

	loadGame.PlayerList = append(loadGame.PlayerList, playerName)

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

	if len(loadGame.PlayerList) < 2 {
		return nil, errors.New("cannot start game: two players are needed to start game")
	}

	if loadGame.GameOver {
		return nil, errors.New("cannot start game: this game is already over")
	}

	if loadGame.GameStarted {
		return loadGame, nil
	}

	for player := range loadGame.Players {
		var randomStartingTiles []string
		for i := 0; i < 7; i++ {
			randomTile := getRandomTile(loadGame.AvailableLetters)
			if randomTile != " " {
				randomStartingTiles = append(randomStartingTiles, randomTile)
			}
		}
		if copyPlayer, ok := loadGame.Players[player]; ok {
			copyPlayer.Hand = randomStartingTiles
			loadGame.Players[player] = copyPlayer
		}
	}

	loadGame.CurrentPlayer = loadGame.PlayerList[0]

	loadGame.GameStarted = true

	app.DatabaseClient.UpdateGameToDB(gameID, *loadGame)

	return loadGame, nil
}

// Load Game by GameID
func (app *App) GetGameById(gameID string) (*Game, error) {
	exists, _ := app.DatabaseClient.CheckGameExists(gameID)
	if !(*exists) {
		return nil, nil
	}
	loadedGame, err := app.DatabaseClient.GetGameByGameID(gameID)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return nil, err
	}

	fmt.Println(loadedGame)

	return loadedGame, nil
}

func (app *App) UpdateGameState(gameID string, playerMove []Move, playerName string) (*Game, error) {
	loadedGame, err := app.GetGameById(gameID)
	if err != nil {
		return nil, err
	}

	if playerName != loadedGame.CurrentPlayer {
		return nil, errors.New("wait! not your turn")
	}

	// update the board once every move is validated and get random tiles to replace tiles used
	for _, move := range playerMove {
		if app.ValidateMove(move, playerName, gameID) {
			loadedGame = updateBoardAndHand(*loadedGame, move, playerName)
		} else {
			return nil, errors.New("invalid move")
		}
	}

	for row := 0; row < 15; row++ {
		for column := 0; column < 15; column++ {
			fmt.Print(loadedGame.Board[row][column], " ")
		}
		fmt.Print("\n")
	}

	// Get score for entered word
	wordScore, err := app.LanguageClient.scoring(*loadedGame, playerMove)
	if err != nil {
		return nil, err
	}

	loadedGame.Players, err = updateScore(wordScore, loadedGame.Players, playerName)
	if err != nil {
		return nil, err
	}

	loadedGame.TotalMoves += 1

	loadedGame.CurrentPlayer = loadedGame.PlayerList[loadedGame.TotalMoves%2]

	// check if game is over
	if !checkPlayerHand(loadedGame.Players) || !checkGameBag(loadedGame.AvailableLetters) {
		loadedGame.GameOver = true
		if loadedGame.Players[loadedGame.PlayerList[0]].Score > loadedGame.Players[loadedGame.PlayerList[1]].Score {
			loadedGame.Winner = loadedGame.PlayerList[0]
		} else {
			loadedGame.Winner = loadedGame.PlayerList[1]
		}
	}

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	return loadedGame, nil
}

// refreshHand refreshes the hand of a player in a game by returning their current tiles to the bag and drawing new tiles from the bag.
// It takes in the loadedGame object and the name of the player whose hand needs to be refreshed.
// It returns a pointer to the updated loadedGame object.
func (app *App) RefreshHand(gameID string, playerName string) (*[]string, error) {
	loadedGame, err := app.GetGameById(gameID)
	newTiles := []string{}
	if err != nil {
		return nil, err
	}

	if loadedGame.CurrentPlayer != playerName {
		return nil, errors.New("wait! not your turn")
	}

	for index, letter := range loadedGame.Players[playerName].Hand {
		returnTilesToBag(*loadedGame, []Move{{Letter: letter}})
		newTile := getRandomTile(loadedGame.AvailableLetters)
		loadedGame.Players[playerName].Hand[index] = newTile
		newTiles = append(newTiles, newTile)
	}

	loadedGame.TotalMoves += 1

	loadedGame.CurrentPlayer = loadedGame.PlayerList[loadedGame.TotalMoves%2]

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	return &newTiles, nil
}

func (app *App) SkipTurn(gameID string, playerName string) (*string, error) {
	loadedGame, err := app.GetGameById(gameID)
	returnMsg := "Your turn is skipped!"
	if err != nil {
		return nil, err
	}

	if loadedGame.CurrentPlayer != playerName {
		return nil, errors.New("wait! not your turn")
	}

	loadedGame.TotalMoves += 1

	loadedGame.CurrentPlayer = loadedGame.PlayerList[loadedGame.TotalMoves%2]

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	return &returnMsg, nil
}

func (app *App) ResignGame(gameID string, playerName string) (*string, error) {
	loadedGame, err := app.GetGameById(gameID)
	winner := ""
	if err != nil {
		return nil, err
	}

	if loadedGame.CurrentPlayer != playerName {
		return nil, errors.New("wait! not your turn")
	}

	loadedGame.TotalMoves += 1

	loadedGame.GameOver = true

	loadedGame.CurrentPlayer = loadedGame.PlayerList[loadedGame.TotalMoves%2]

	for _, player := range loadedGame.PlayerList {
		if player != playerName {
			winner = player
		}
	}

	loadedGame.Winner = winner

	returnMsg := fmt.Sprintf("Player %s resigned. Player %s is the winner!", playerName, winner)

	// update game on database
	app.DatabaseClient.UpdateGameToDB(gameID, *loadedGame)

	return &returnMsg, nil
}

func generateNewGameID() string {
	gameID := uniuri.NewLen(6)
	return gameID
}

// get a random tile from the game tile bag
func getRandomTile(availableLetters map[string]int) string {
	var keys []string
	// get list of tiles that are available
	for k := range availableLetters {
		if availableLetters[k] > 0 {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return " "
	}
	// in golang, iteration order is not specified and is not guaranteed to be the same from one iteration to the next
	// this will therefore return a random value
	availableLetters[keys[0]] -= 1
	return keys[0]
}

// Update the Board with player's move
func updateBoardAndHand(loadedGame Game, playerMove Move, playerName string) *Game {
	// update board state
	loadedGame.Board[playerMove.Col][playerMove.Row] = playerMove.Letter
	index := 0
	for _, iterateLetter := range loadedGame.Players[playerName].Hand {
		randomTile := getRandomTile(loadedGame.AvailableLetters)
		if iterateLetter == playerMove.Letter {
			if randomTile != " " {
				loadedGame.Players[playerName].Hand[index] = randomTile
			}
			break
		}
		index++
	}

	return &loadedGame
}

// returnTilesToBag returns the tiles used in the player's move back to the game's available letters.
// loadedGame is the game being played.
// playerMove is a slice of moves made by the player.
func returnTilesToBag(loadedGame Game, playerMove []Move) {
	for _, move := range playerMove {
		loadedGame.AvailableLetters[move.Letter] += 1
	}
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

func checkPlayerHand(currPlayers map[string]PlayerInfo) bool {
	// check if any player hands are empty
	for _, currPlayerInfo := range currPlayers {
		if len(currPlayerInfo.Hand) == 0 {
			return false
		}
	}
	return true
}

func checkGameBag(availableLetters map[string]int) bool {
	keys := []string{}
	// check if any tiles are left in the tile bag
	for k := range availableLetters {
		if availableLetters[k] > 0 {
			keys = append(keys, k)
		}
	}

	return len(keys) > 0
}
