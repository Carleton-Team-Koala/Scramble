package models

import (
	"errors"
	"fmt"
)

func (c *LanguageClient) scoring(activeGame Game, newTiles []Move) (int, error) {
	score := 0
	word, wordTwo, wordThree := "", "", ""
	up, down, left, right := WordScore{}, WordScore{}, WordScore{}, WordScore{}

	for i := 0; i < len(newTiles); i++ {
		word = word + newTiles[i].Letter

		print("word so far: " + word + "\n")

		score += c.GetLetterScore(newTiles[i].Letter)

		var x = int(newTiles[i].XLoc)
		var y = int(newTiles[i].YLoc)

		var boolXDir = checkTwoDirectionsX(activeGame, x, y)
		var boolYDir = checkTwoDirectionsY(activeGame, x, y)

		fmt.Printf("%t\n", boolXDir)
		fmt.Printf("%t\n", boolYDir)

		if boolXDir {
			wordTwo = pullLeft(activeGame, x, y) + activeGame.Board[x][y] + pullRight(activeGame, x, y)
		} else {
			up = c.checkUp(x, y, activeGame, newTiles[i].Letter, "", 0)
			down = c.checkDown(x, y, activeGame, newTiles[i].Letter, "", 0)

			if up.Valid {
				score += up.Score
			}
			if down.Valid {
				score += down.Score
			}

			fmt.Printf("up:%t\n", up.Valid)
			fmt.Printf("down:%t\n", down.Valid)

			if !up.Valid || !down.Valid {
				return 0, errors.New("invalid Words up/down")
			}
		}

		if boolYDir {
			wordThree = pullUp(activeGame, x, y) + activeGame.Board[x][y] + pullDown(activeGame, x, y)
		} else {
			left = c.checkLeft(x, y, activeGame, newTiles[i].Letter, "", 0)
			right = c.checkRight(x, y, activeGame, newTiles[i].Letter, "", 0)

			if left.Valid {
				score += left.Score
			}
			if right.Valid {
				score += right.Score
			}

			if !right.Valid || !left.Valid {
				return 0, errors.New("Invalid Words left/right")
			}
		}

	}

	if c.CheckValidWord(word) {
		if c.CheckValidWord(wordTwo) || wordTwo == "" {
			if c.CheckValidWord(wordThree) || wordThree == "" {
				return score, nil
			}
			return 0, errors.New("Invalid Words:" + wordThree)
		}
		return 0, errors.New("Invalid Words:" + wordTwo)
	}

	return 0, errors.New("Score not calculated")
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
	if activeGame.Board[x-1][y] != "" && activeGame.Board[x+1][y] != "" {
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

func (c *LanguageClient) endOfWord(scoreValue int, word string) WordScore {
	var foo WordScore
	var validWord = c.CheckValidWord(word)

	print("end of word check: " + word + "\n")
	if word == "" {
		foo = WordScore{
			Valid: true,
			Score: 0,
		}
	} else {
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
	}
	return foo

}

func (c *LanguageClient) checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x > 0 {
		if activeGame.Board[x-1][y] != "" {
			newScore := scoreValue + c.GetLetterScore(activeGame.Board[x-1][y])
			return c.checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, newScore)
		}
		return c.endOfWord(scoreValue, activeGame.Board[x-1][y]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func (c *LanguageClient) checkRight(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x < 16 {
		if activeGame.Board[x+1][y] != "" {
			newScore := scoreValue + c.GetLetterScore(activeGame.Board[x+1][y])
			return c.checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x+1][y]+wordSoFar, newScore)
		}
		return c.endOfWord(scoreValue, activeGame.Board[x+1][y]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func (c *LanguageClient) checkUp(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y > 0 {
		if activeGame.Board[x][y-1] != "" {
			newScore := scoreValue + c.GetLetterScore(activeGame.Board[x][y-1])
			return c.checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x][y-1]+wordSoFar, newScore)
		}
		return c.endOfWord(scoreValue, activeGame.Board[x][y-1]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func (c *LanguageClient) checkDown(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y < 16 {
		if activeGame.Board[x][y+1] != "" {
			newScore := scoreValue + c.GetLetterScore(activeGame.Board[x][y+1])
			return c.checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x][y+1]+wordSoFar, newScore)
		}
		return c.endOfWord(scoreValue, activeGame.Board[x][y+1]+wordSoFar)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}
