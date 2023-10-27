package models

import "fmt"

// functions used for validating user move

// TODO: Add more validating functions(checkLeft, checkRight, existingWord etc)
func (app *App) ValidateMove(playerMove Move, playerName string, gameID string) bool {
	loadGame, err := app.GetGameById(gameID)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return false
	}

	// check if letter is available
	if checkLetterAvailability(playerMove.Letter, loadGame.Players[playerName].Hand) != nil {
		fmt.Println(fmt.Errorf("ValidateMove:checkLetterAvailability: %v", err))
		return false
	}

	// check if cell location is valid
	if checkLocation(playerMove.XLoc, playerMove.YLoc, loadGame.Board) != nil {
		fmt.Println(fmt.Errorf("ValidateMove:checkLocation: %v", err))
		return false
	}
	return true
}

func checkLetterAvailability(letter string, availableLetters []string) error {
	for _, tile := range availableLetters {
		if tile == letter {
			return nil
		}
	}
	return fmt.Errorf("letter unavailable")
}

func checkLocation(xLoc int, yLoc int, gameBoard [15][15]string) error {
	if 0 > xLoc || 15 < xLoc || 0 > yLoc || 15 < yLoc {
		return fmt.Errorf("invalid cell location")
	}

	if gameBoard[xLoc][yLoc] != "" {
		return fmt.Errorf("cell already in use")
	}

	return nil
}

