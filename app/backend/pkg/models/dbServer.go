package models

import (
	"encoding/json"
	"fmt"
)

// add new game to database
func (c *DatabaseClient) AddNewGameToDB(newGame Game) error {
	newBoard, newLetters, newPlayers, newPlayerList, err := jsonifyGameField(newGame)
	if err != nil {
		return fmt.Errorf("addNewGameToDB: %v", err)
	}
	_, err = c.database.Exec("INSERT INTO games (GameID, Board, LetterDistribution, Players, CurrentPlayer, PlayerList, TotalMoves, GameStarted) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", newGame.GameID, newBoard, newLetters, newPlayers, newGame.CurrentPlayer, newPlayerList, newGame.TotalMoves, newGame.GameStarted)
	if err != nil {
		return fmt.Errorf("addNewGameToDB: %v", err)
	}
	fmt.Println("Done adding game to database!")
	return nil
}

// get game from database using game ID
func (c *DatabaseClient) GetGameByGameID(gameID string) (*Game, error) {
	selectedGame := Game{}
	var jsonBoard []byte
	var jsonLetters []byte
	var jsonPlayers []byte
	var jsonPlayerList []byte

	selectedRow := c.database.QueryRow("SELECT GameID, Board, LetterDistribution, Players, CurrentPlayer, PlayerList, TotalMoves, GameStarted FROM games WHERE \"gameid\"=$1", gameID)
	err := selectedRow.Scan(&selectedGame.GameID, &jsonBoard, &jsonLetters, &jsonPlayers, &selectedGame.CurrentPlayer, &jsonPlayerList, &selectedGame.TotalMoves, &selectedGame.GameStarted)
	if err != nil {
		return nil, fmt.Errorf("getGameByGameID: %v", err)
	}
	selectedGame.Board, selectedGame.AvailableLetters, selectedGame.Players, selectedGame.PlayerList, err = unJsonifyGameField(jsonBoard, jsonLetters, jsonPlayers, jsonPlayerList)
	if err != nil {
		return nil, fmt.Errorf("getGameByGameID: %v", err)
	}

	fmt.Println("Done selecting game by gameID!")
	return &selectedGame, nil
}

// get game from database using game ID
func (c *DatabaseClient) UpdateGameToDB(gameID string, updatedGame Game) error {
	newBoard, newLetters, newPlayers, newPlayerList, err := jsonifyGameField(updatedGame)
	if err != nil {
		return fmt.Errorf("updateGameToDB: %v", err)
	}

	_, err = c.database.Exec("UPDATE games SET board=$1, letterdistribution=$2, players=$3, currentplayer=$4, playerlist=$5, totalmoves=$6, gamestarted=$7 WHERE \"gameid\"=$8", newBoard, newLetters, newPlayers, updatedGame.CurrentPlayer, newPlayerList, updatedGame.TotalMoves, updatedGame.GameStarted, gameID)
	if err != nil {
		return fmt.Errorf("updateGameToDB: %v", err)
	}

	fmt.Printf("Done updating game on database: %s\n", gameID)
	return nil
}

// get game from database using game ID
func (c *DatabaseClient) CheckGameExists(gameID string) (*bool, error) {
	var exists bool
	selectedRow := c.database.QueryRow("SELECT count(1) > 0 FROM games WHERE \"gameid\" = $1", gameID)
	err := selectedRow.Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("checkGameExists: %v", err)
	}

	fmt.Println("Done checking game by gameID!")
	return &exists, nil
}

// jsonify game fields
func jsonifyGameField(game Game) ([]byte, []byte, []byte, []byte, error) {
	jsonBoard, err := json.Marshal(game.Board)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	jsonLetters, err := json.Marshal(game.AvailableLetters)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	jsonPlayers, err := json.Marshal(game.Players)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	jsonPlayerList, err := json.Marshal(game.PlayerList)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	return jsonBoard, jsonLetters, jsonPlayers, jsonPlayerList, nil
}

// unjsonify game fields
func unJsonifyGameField(selectedBoard []byte, selectedLetters []byte, selectedPlayers []byte, selectedPlayerList []byte) ([15][15]string, map[string]int, map[string]PlayerInfo, []string, error) {
	var board [15][15]string
	if err := json.Unmarshal([]byte(selectedBoard), &board); err != nil {
		return [15][15]string{}, nil, nil, nil, fmt.Errorf("failed to unmarshal board: %w", err)
	}
	var letters map[string]int
	if err := json.Unmarshal([]byte(selectedLetters), &letters); err != nil {
		return [15][15]string{}, nil, nil, nil, fmt.Errorf("failed to unmarshal letters: %w", err)
	}
	var players map[string]PlayerInfo
	if err := json.Unmarshal([]byte(selectedPlayers), &players); err != nil {
		return [15][15]string{}, nil, nil, nil, fmt.Errorf("failed to unmarshal players: %w", err)
	}
	var playerList []string
	if err := json.Unmarshal([]byte(selectedPlayerList), &playerList); err != nil {
		return [15][15]string{}, nil, nil, nil, fmt.Errorf("failed to unmarshal player list: %w", err)
	}
	return board, letters, players, playerList, nil
}
