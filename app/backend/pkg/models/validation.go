package models

import "fmt"

// function used for validating user move
func (app *App) ValidateMove(playerMove Move, playerName string, gameID string) bool {
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		return false
	}

	// check if letter is available
	if checkLetterAvailability(playerMove.Letter, loadGame.Players[playerName].Hand) != nil {
		return false
	}

	// check if cell location is valid
	if checkLocation(playerMove.Col, playerMove.Row, loadGame.Board) != nil {
		return false
	}
	return true
}

// check whether letter is available
func checkLetterAvailability(letter string, availableLetters []string) error {
	for _, tile := range availableLetters {
		if tile == letter {
			return nil
		}
	}
	return fmt.Errorf("letter unavailable")
}

// check location is valid
func checkLocation(xLoc int, yLoc int, gameBoard [15][15]string) error {
	if 0 > xLoc || 15 < xLoc || 0 > yLoc || 15 < yLoc {
		return fmt.Errorf("invalid cell location")
	}

	if gameBoard[xLoc][yLoc] != "" {
		return fmt.Errorf("cell already in use")
	}

	return nil
}
