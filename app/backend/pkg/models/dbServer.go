package models

import (
	"fmt"
)

// TODO: modify this function
// connect to languages to get the list of available tiles when creating game
func (c *DatabaseClient) AddNewGameToDB(newGame Game) error {
	_, err := c.database.Exec("INSERT INTO games (gameID, board, letterdistribution, players) VALUES (?, scan_rows(?), ?, ?)", newGame.GameID, newGame.Board, newGame.AvailableLetters, newGame.Players)
	if err != nil {
		return fmt.Errorf("addNewGameToDB: %v", err)
	}
	fmt.Println("done!")
	return nil
}
