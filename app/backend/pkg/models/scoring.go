package models

func scoring(activeGame Game, newTiles string) {
	var word = ""
	for i := 0; i < len(newTiles); i++ {
		word.append(newTiles[i])

	}
}

func checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x >= 0 {
		if activeGame.Board[x-1][y] != "" {
			return checkLeft(x-1, y, activeGame, newTiles, activeGame.Board[x-1][y]+wordSoFar, scoreValue)
		}
		var foo WordScore
		// TODO: connect to langauges api for scoring/ word validation
		var validWord bool
		if validWord {
			foo := WordScore{
				Valid: true,
				Score: scoreValue + 0,
			}
		} else {
			foo := WordScore{
				Valid: false,
				Score: scoreValue + 0,
			}
		}
		return foo

	}
	foo := WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

/*

func checkRight(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if x > 0 {
		if activeGame.Board[x+1][y] != ""{
			return checkLeft(x+1, y, activactiveGame, newTiles, activeGame.Board[x-1][y] + wordSoFar, scoreValue)
		}

			// TODO: connect to langauges api for scoring/ word validation
			var validWord bool
			if validWord{
				foo := WordScore{
					Valid: true,
					Score: scoreValue + 0,
				}
			} else {
				foo := WordScore{
					Valid: false,
					Score: scoreValue + 0,
				}
			}
			return foo

	}
	foo := WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}

func checkUp(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y > 0 {
		if activeGame.Board[x][y-1] != ""{
			return checkLeft(x, y-1, activactiveGame, newTiles, activeGame.Board[x-1][y] + wordSoFar, scoreValue)
		}

			// TODO: connect to langauges api for scoring/ word validation
			var validWord bool
			if validWord{
				foo := WordScore{
					Valid: true,
					Score: scoreValue + 0,
				}
			} else {
				foo := WordScore{
					Valid: false,
					Score: scoreValue + 0,
				}
			}
			return foo

	}
}


func checkDown(x int, y int, activeGame Game, newTiles string, wordSoFar string, scoreValue int) WordScore {
	if y > 0 {
		if activeGame.Board[x][y+1] != ""{
			return checkLeft(x. y+1, activactiveGame, newTiles, activeGame.Board[x-1][y] + wordSoFar, scoreValue)
		}
			// TODO: connect to langauges api for scoring/ word validation
			var validWord bool
			if validWord{
				foo := WordScore{
					Valid: true,
					Score: scoreValue + 0,
				}
			} else {
				foo := WordScore{
					Valid: false,
					Score: scoreValue + 0,
				}
			}
			return foo

	}
	foo := WordScore{
		Valid: true,
		Score: 0,
	}
	return foo
}


*/
