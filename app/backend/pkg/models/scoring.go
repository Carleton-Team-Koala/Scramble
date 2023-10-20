package models

func scoring(activeGame Game, newTiles string) {
	var word = ""
	for i := 0; i < len(newTiles); i++ {
		word = word + string(newTiles[i])
	}
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

func checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
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

func checkRight(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
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

func checkUp(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
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

func checkDown(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
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
