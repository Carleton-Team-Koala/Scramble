package models

import (
	"errors"
)

func scoring(activeGame Game, newTiles string) (int, error) {
	score := 0
	word := ""
	for i := 0; i < len(newTiles); i++ {
		word = word + string(newTiles[i].letter)

		score += GetLetterScore(newTiles[i].letter)

		var x = int(newTiles.x)
		var y = int(newTiles.y)

		var boolXDir = checkTwoDirectionsX(activeGame, x, y)
		var boolYDir = checkTwoDirectionsX(activeGame, x, y)

		if boolXDir {
			wordTwo := pullLeft(activeGame, x, y) + game.Board[x][y] + pullRight(activeGame, x, y)
		} else {
			up := checkUp(x, y, activeGame, newTiles, "", 0)
			down := checkDown(x, y, activeGame, newTiles, "", 0)

			if up.Valid {
				score += up.Score
			}
			if down.Valid {
				score += down.Score
			}

			if !up.Valid || !down.Valid {
				return 0, errors.New("Invalid Words up/down")
			}
		}

		if boolYDir {
			wordThree := pullUp(activeGame, x, y) + game.Board[x][y] + pullDown(activeGame, x, y)
		} else {
			left := checkLeft(x, y, activeGame, newTiles, "", 0)
			right := checkRight(x, y, activeGame, newTiles, "", 0)

			if left.Valid {
				score += up.Score
			}
			if right.Valid {
				score += down.Score
			}

			if !right.Valid || !left.Valid {
				return 0, errors.New("Invalid Words left/right")
			}
		}

	}

	if CheckValidWord(word) {
		if CheckValidWord(wordTwo) || wordTwo == "" {
			if CheckValidWord(wordThree) || wordThree == "" {
				return score, nil
			}
			return 0, errors.New("Invalid Words:" + wordThree)
		}
		return 0, errors.New("Invalid Words:" + wordTwo)
	}

}

func pullUp(game Game, x int, y int) string {
	if game.Board[x][y-1] == "" {
		return ""
	}
	return pullUp(game, x, y-1) + game.Board[x][y-1]
}

func pullDown(game Game, x int, y int) string {
	if game.Board[x][y+1] == "" {
		return ""
	}
	return game.Board[x][y+1] + pullDown(game, x, y+1)
}

func pullLeft(game Game, x int, y int) string {
	if game.Board[x-1][y] == "" {
		return ""
	}
	return pullUp(game, x-1, y) + game.Board[x-1][y]
}

func pullRight(game Game, x int, y int) string {
	if game.Board[x+1][y] == "" {
		return ""
	}
	return game.Board[x+1][y] + pullDown(game, x+1, y)
}

func checkTwoDirectionsX(activeGame Game, x int, y int) bool {
	if x <= 0 || x >= 15 {
		return false
	}
	if activeGame.Board[x-1][y] != nil && activeGame.Board[x+1][y] != nil {
		return true
	}
	return false
}

func checkTwoDirectionsY(activeGame Game, x int, y int) bool {
	if y <= 0 || y >= 15 {
		return false
	}
	if activeGame.Board[x][y-1] != "" && activeGame.Board[x][y+1] != "" {
		return true
	}
	return false
}

func endOfWord(scoreValue int, word string) WordScore {
	var foo WordScore
	var validWord = CheckValidWord(word)
	if validWord {
		foo = WordScore{
			Valid: true,
			Score: scoreValue + 0,
		}
	} else {
		foo = WordScore{
			Valid: false,
			Score: 0,
		}
	}
	return foo

}

func checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x > 0 {
		if activeGame.Board[x-1][y] != "" {
			newScore := scoreValue + GetLetterScore(activeGame.Board[x-1][y])
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, newScore)
		}
		return endOfWord(scoreValue, activeGame.Board[x-1][y]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkRight(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x < 16 {
		if activeGame.Board[x+1][y] != "" {
			newScore := scoreValue + GetLetterScore(activeGame.Board[x+1][y])
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x+1][y]+wordSoFar, newScore)
		}
		return endOfWord(scoreValue, activeGame.Board[x+1][y]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkUp(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y > 0 {
		if activeGame.Board[x][y-1] != "" {
			newScore := scoreValue + GetLetterScore(activeGame.Board[x][y-1])
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x][y-1]+wordSoFar, newScore)
		}
		return endOfWord(scoreValue, activeGame.Board[x][y-1]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkDown(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y < 15 {
		if activeGame.Board[x][y+1] != "" {
			newScore := scoreValue + GetLetterScore(activeGame.Board[x][y+1])
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x][y+1]+wordSoFar, newScore)
		}
		return endOfWord(scoreValue, activeGame.Board[x][y+1]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}
