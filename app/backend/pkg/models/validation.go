package models

import "fmt"

// functions used for validating user move

// TODO: Add more validating functions(checkLeft, checkRight, existingWord etc)
func ValidateMove(playerMove Move, gameID string) bool {
	loadGame := GetGameById(gameID)

	// check if game id does not exist
	if checkGameExists(gameID) != nil {
		return false
	}

	// check if letter is available
	if checkLetterAvailability(playerMove.Letter, loadGame.AvailableLetters) != nil {
		return false
	}

	// check if cell location is valid
	if checkLocation(playerMove.XLoc, playerMove.YLoc, loadGame.Board) != nil {
		return false
	}
	return true
}

func checkLetterAvailability(letter string, availableLetters map[string]int) error {
	if availableLetters[letter] > 0 {
		return nil
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

// check if gameID exists. If not, return error
func checkGameExists(gameID string) error {
	_, exists := GameList[gameID]
	if !exists {
		return fmt.Errorf("Game ID not found")
	}
	return nil
}
