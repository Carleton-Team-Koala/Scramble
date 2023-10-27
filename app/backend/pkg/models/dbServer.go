package models

import (
	"encoding/json"
	"fmt"
)

// add new game to database
func (c *DatabaseClient) AddNewGameToDB(newGame Game) error {
	newBoard, newLetters, newPlayers, err := jsonifyGameField(newGame)
	if err != nil {
		return fmt.Errorf("addNewGameToDB: %v", err)
	}
	_, err = c.database.Exec("INSERT INTO games (GameID, Board, LetterDistribution, Players) VALUES($1, $2, $3, $4)", newGame.GameID, newBoard, newLetters, newPlayers)
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

	selectedRow := c.database.QueryRow("SELECT GameID, Board, LetterDistribution, Players FROM games WHERE \"gameid\"=$1", gameID)
	err := selectedRow.Scan(&selectedGame.GameID, &jsonBoard, &jsonLetters, &jsonPlayers)
	if err != nil {
		return nil, fmt.Errorf("getGameByGameID: %v", err)
	}
	selectedGame.Board, selectedGame.AvailableLetters, selectedGame.Players, err = unJsonifyGameField(jsonBoard, jsonLetters, jsonPlayers)
	if err != nil {
		return nil, fmt.Errorf("getGameByGameID: %v", err)
	}

	fmt.Println("Done selecting game by gameID!")
	return &selectedGame, nil
}

// get game from database using game ID
func (c *DatabaseClient) UpdateGameToDB(gameID string, updatedGame Game) (error) {
	newBoard, newLetters, newPlayers, err := jsonifyGameField(updatedGame)
	if err != nil {
		return fmt.Errorf("updateGameToDB: %v", err)
	}

	_, err = c.database.Exec("UPDATE games SET board=$1, letterdistribution=$2, players=$3 WHERE \"gameid\"=$4", newBoard, newLetters, newPlayers, gameID)
	if err != nil {
		return fmt.Errorf("updateGameToDB: %v", err)
	}
	
	fmt.Println("Done updating game %s on database", gameID)
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
func jsonifyGameField(game Game) ([]byte, []byte, []byte, error) {
	jsonBoard, err := json.Marshal(game.Board)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	jsonLetters, err := json.Marshal(game.AvailableLetters)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	jsonPlayers, err := json.Marshal(game.Players)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("jsonifyGameField: %v", err)
	}

	return jsonBoard, jsonLetters, jsonPlayers, nil
}

// unjsonify game fields
func unJsonifyGameField(selectedBoard []byte, selectedLetters []byte, selectedPlayers []byte) ([15][15]string, map[string]int, map[string]PlayerInfo, error) {
	var board [15][15]string
	if err := json.Unmarshal([]byte(selectedBoard), &board); err != nil {
		return [15][15]string{}, nil, nil, fmt.Errorf("failed to unmarshal board: %w", err)
	}
	var letters map[string]int
	if err := json.Unmarshal([]byte(selectedLetters), &letters); err != nil {
		return [15][15]string{}, nil, nil, fmt.Errorf("failed to unmarshal letters: %w", err)
	}
	var players map[string]PlayerInfo
	if err := json.Unmarshal([]byte(selectedPlayers), &players); err != nil {
		return [15][15]string{}, nil, nil, fmt.Errorf("failed to unmarshal players: %w", err)
	}
	return board, letters, players, nil
}
