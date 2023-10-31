package models

import "container/list"

func scoring(activeGame Game, newTiles string)  {
	score := 0
	word = ""
	for i := 0; i < len(newTiles); i++ {
		word = word + string(newTiles[i].letter)

		var x = int(newTiles.x)
		var y = int(newTiles.y)

		var boolXDir = checkTwoDirectionsX(activeGame, x, y)
		var boolYDir = checkTwoDirectionsX(activeGame, x, y)

		if boolXDir {
			word = pullLeft(activeGame, x, y) + pullRight(activeGame, x, y)
		} else {
			score += checkUp(x, y, activeGame, newTiles, "", 0)
			score += checkDown(x, y, activeGame, newTiles, "", 0)
		}

		if boolYDir {
			word = 	pullUp(activeGame, x, y) + pullDown(activeGame, x, y)
		} else {
			score += checkUp(x, y, activeGame, newTiles, "", 0)
			score += checkDown(x, y, activeGame, newTiles, "", 0)
		}

	}

	if CheckValidWord(word){
		
	}



	return score
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
	return game.Board[x][y+1]+ pullDown(game, x, y+1)
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
	return game.Board[x+1][y]+ pullDown(game, x+1, y)
}

func checkTwoDirectionsX(activeGame Game, x int y int) tf bool{
	if (x <= 0 || x >= 15){
		return false
	}
	if (activeGame.Board[x-1][y] != nil && activeGame.Board[x+1][y] != nil){
		return true
	}
	return false
}

func checkTwoDirectionsY(activeGame Game, x int y int) tf bool{
	if (y <= 0 || y >= 15){
		return false
	}
	if (activeGame.Board[x][y-1] != nil && activeGame.Board[x][y+1] != nil){
		return true
	}
	return false
}

func endOfWord(scoreValue int) WordScore {
	var foo WordScore
	// TODO: connect to langauges api for scoring/ word validation
	var validWord bool
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

func checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore int {
	if x > 0 {
		if activeGame.Board[x-1][y] != "" {
			//TODO - score individual tile to add to scoreValue
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, scoreValue)
		}
		return endOfWord(scoreValue)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkRight(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore  int{
	if x < 16 {
		if activeGame.Board[x+1][y] != "" {
			//TODO - score individual tile to add to scoreValue
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, scoreValue)
		}
		return endOfWord(scoreValue)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkUp(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore int{
	if y > 0 {
		if activeGame.Board[x][y-1] != "" {
			//TODO - score individual tile to add to scoreValue
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, scoreValue)
		}
		return endOfWord(scoreValue)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkDown(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore int {
	if y < 15 {
		if activeGame.Board[x][y+1] != "" {
			//TODO - score individual tile to add to scoreValue
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, scoreValue)
		}
		return endOfWord(scoreValue)
	}

	var foo = WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}
