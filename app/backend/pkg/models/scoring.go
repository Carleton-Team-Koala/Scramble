package models

// scoring.go contains the scoring logic for the game of Scramble.
import (
	"errors"
	"reflect"
	"sort"
)

type MoveSlice []Move

func (m MoveSlice) Len() int {
	return len(m)
}

// Less compares two moves in the MoveSlice based on their column and row values.
// It returns true if the move at index i should be considered "less" than the move at index j.
// Moves are considered "less" if their column values are smaller, or if their column values are equal
// and their row values are smaller.
func (m MoveSlice) Less(i, j int) bool {
	// First compare by Col.
	if m[i].Col != m[j].Col {
		return m[i].Col < m[j].Col
	}
	// If Col is the same, compare by Row.
	return m[i].Row < m[j].Row
}

func (m MoveSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// scoring calculates the score for a given move in a game.
// It takes the activeGame, which represents the current state of the game,
// and newTiles, which represents the tiles being placed in the move.
// It returns the calculated score and an error if any invalid conditions are encountered.
func (c *LanguageClient) scoring(activeGame Game, newTiles MoveSlice) (int, error) {
	isFirstMove := isFirstMove(activeGame, newTiles)

	if isFirstMove != nil {
		return 0, isFirstMove
	}

	score := 0
	setOfWords := []string{}
	scoreModifier := [15][2]int{}
	OriginalWord := ""
	scoreAggregateModifier := 1
	OGWordScore := 0

	sort.Sort(newTiles)

	adjacentToPlacedTile := TestAdjacentToPlacedTile(activeGame, newTiles)
	if !adjacentToPlacedTile {
		return 0, errors.New("at least one new tile must be adjacent to an already placed tile")
	}

	isSequentialInX, isSequentialInY := checkSequential(newTiles, activeGame)

	// check for bisecting tile
	if !isSequentialInX && !isSequentialInY {
		var bisectingTile *Move
		for i := 0; i < len(newTiles); i++ {
			for row := 0; row < len(activeGame.Board); row++ {
				for col := 0; col < len(activeGame.Board[row]); col++ {
					if activeGame.Board[row][col] == string(newTiles[i].Letter) {
						bisectingTile = &newTiles[i]
						break
					}
				}
			}
		}
		if bisectingTile != nil {
			// check if new tiles are placed in a single row or column on either side of the bisecting tile
			isSequentialInX = true
			isSequentialInY = true
			for _, tile := range newTiles {
				if tile.Row != bisectingTile.Row {
					isSequentialInX = false
				}
				if tile.Col != bisectingTile.Col {
					isSequentialInY = false
				}
			}
			if !isSequentialInX && !isSequentialInY {
				return 0, errors.New("new tiles must be placed in a single row or column")
			}
		} else {
			return 0, errors.New("new tiles must be placed in a single row or column")
		}
	}

	isAdjacent := TestAdjacentToPlacedTile(activeGame, newTiles)

	if !isAdjacent {
		return 0, errors.New("at least one new tile must be adjacent to an already placed tile")
	}
	// }

	// check if all new tiles are in a line either horizontally or vertically
	for i := 0; i < len(newTiles); i++ {

		var x = int(newTiles[i].Col)
		var y = int(newTiles[i].Row)

		OriginalWord += string(newTiles[i].Letter)
		scoreModifier[i] = [2]int{x, y}

		// recursively get all the possible words
		leftAndRightWord := pullLeft(activeGame, x, y) + activeGame.Board[x][y] + pullRight(activeGame, x, y)
		upAndDownWord := pullUp(activeGame, x, y) + activeGame.Board[x][y] + pullDown(activeGame, x, y)

		if (!c.CheckValidWord(leftAndRightWord) && len(leftAndRightWord) > 1) || (!c.CheckValidWord(upAndDownWord) && len(upAndDownWord) > 1) {
			return 0, errors.New("this is an invalid word: " + leftAndRightWord + " or " + upAndDownWord)
		}

		if len(leftAndRightWord) < 2 && len(upAndDownWord) < 2 && len(newTiles) == 1 {
			return 0, errors.New("words must be longer than a single letter")
		}

		// then append to the list of words that would count towards the scores
		if !checkWordExists(setOfWords, leftAndRightWord) && len(leftAndRightWord) > 1 {
			setOfWords = append(setOfWords, leftAndRightWord)
		}

		if !checkWordExists(setOfWords, upAndDownWord) && len(upAndDownWord) > 1 {
			setOfWords = append(setOfWords, upAndDownWord)
		}
	}

	// OGWordCalculated := false
	// calculate the score
	for _, word := range setOfWords {
		OGWordScore = 0
		scoreAggregateModifier = 1
		if c.CheckValidWord(word) {
			//scoring piece by piece, adding allowances for score modifiers
			i := 0
			indexOfSM := 0
			for i < len(word) {
				// valueOfModifier is either "dl", "tl", "dw", "tw", or "na" for double letter, triple letter, double word, triple word, or no modifier, respectively.
				valueOfModifier := checkForScoreModifier(scoreModifier[indexOfSM][0], scoreModifier[indexOfSM][1])
				// tileIn is true if the tile at the given position is the same as the current letter in the word.
				tileIn := false
				if activeGame.Board[scoreModifier[indexOfSM][0]][scoreModifier[indexOfSM][1]] == string(word[i]) {
					tileIn = true
				}
				if valueOfModifier == "dl" && tileIn {
					OGWordScore += 2 * (c.GetLetterScore(string(word[i])))
				} else if valueOfModifier == "tl" && tileIn {
					OGWordScore += 3 * (c.GetLetterScore(string(word[i])))
				} else if valueOfModifier == "dw" && tileIn {
					scoreAggregateModifier *= 2
					OGWordScore += (c.GetLetterScore(string(word[i])))
				} else if valueOfModifier == "tw" && tileIn {
					scoreAggregateModifier *= 3
					OGWordScore += (c.GetLetterScore(string(word[i])))
				} else {
					if !tileIn {
						indexOfSM--
					}
					OGWordScore += (c.GetLetterScore(string(word[i])))
				}
				i++
				indexOfSM++
			}

		} else {
			return 0, errors.New("this is an invalid word: " + word)
		}
		score += OGWordScore * scoreAggregateModifier
	}

	// Add 50 points if all 7 tiles are used
	if newTiles.Len() == 7 {
		score += 50
	}

	return score, nil
}

// pullLeft recursively pulls the letters to the left of the given position (x,y) on the game board.
func pullLeft(game Game, x int, y int) string {
	if y <= 0 || game.Board[x][y] == "" {
		return ""
	}
	return pullLeft(game, x, y-1) + game.Board[x][y-1]
}

// pullRight recursively pulls the letters to the right of the given position (x,y) on the game board.
func pullRight(game Game, x int, y int) string {
	if y >= 14 || game.Board[x][y] == "" {
		return ""
	}
	return game.Board[x][y+1] + pullRight(game, x, y+1)
}

// pullUp recursively pulls the letters above the given position (x,y) on the game board.
func pullUp(game Game, x int, y int) string {
	if x <= 0 || game.Board[x][y] == "" {
		return ""
	}
	return pullUp(game, x-1, y) + game.Board[x-1][y]
}

// pullDown recursively pulls the letters below the given position (x,y) on the game board.
func pullDown(game Game, x int, y int) string {

	if x >= 14 || game.Board[x][y] == "" {
		return ""
	}
	return game.Board[x+1][y] + pullDown(game, x+1, y)
}

// checkWordExists checks if a given word exists in a set of words.
func checkWordExists(setOfWords []string, word string) bool {
	for _, eachWord := range setOfWords {
		if eachWord == word {
			return true
		}
	}
	return false
}

// checkForScoreModifier checks if a given position (x,y) on the game board is a score modifier (double word, double letter, triple letter, triple word).
func checkForScoreModifier(x int, y int) string {
	tw := [][2]int{{0, 0}, {0, 7}, {0, 14}, {7, 0}, {7, 14}, {14, 0}, {14, 7}, {14, 14}}
	dw := [][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {10, 10}, {11, 11}, {12, 12}, {13, 13}, {1, 13}, {2, 12}, {3, 11}, {4, 10}, {10, 4}, {11, 3}, {12, 2}, {13, 1}}
	tl := [][2]int{{1, 5}, {1, 9}, {5, 1}, {5, 5}, {5, 9}, {5, 13}, {9, 1}, {9, 5}, {9, 9}, {9, 13}, {13, 5}, {13, 9}}
	dl := [][2]int{{0, 3}, {0, 11}, {2, 6}, {2, 8}, {3, 0}, {3, 7}, {3, 14}, {6, 2}, {6, 6}, {6, 8}, {6, 12}, {7, 3}, {7, 11}, {8, 2}, {8, 6}, {8, 8}, {8, 12}, {11, 0}, {11, 7}, {11, 14}, {12, 6}, {12, 8}, {14, 3}, {14, 11}}

	for i := 0; i < len(tw); i++ {
		tempVal := [2]int{x, y}
		if tempVal == tw[i] {
			tw[i] = [2]int{-1, -1}
			return "tw"
		}
	}

	for i := 0; i < len(dw); i++ {
		tempVal := [2]int{x, y}
		if tempVal == dw[i] {
			dw[i] = [2]int{-1, -1}
			return "dw"
		}
	}

	for i := 0; i < len(tl); i++ {
		tempVal := [2]int{x, y}
		if tempVal == tl[i] {
			tl[i] = [2]int{-1, -1}
			return "tl"
		}
	}

	for i := 0; i < len(dl); i++ {
		tempVal := [2]int{x, y}
		if tempVal == dl[i] {
			dl[i] = [2]int{-1, -1}
			return "dl"
		}
	}

	return "na"
}

// checkSequential verifies if all tiles are in a line either horizontally or vertically.
func checkSequential(tiles MoveSlice, game Game) (bool, bool) {
	isSequentialInX := true
	isSequentialInY := true

	// Check if sequential in X direction (same Row)
	for i := 1; i < len(tiles); i++ {
		if tiles[i].Row != tiles[0].Row {
			isSequentialInY = false
			break
		}
	}

	// If not sequential in X, check if sequential in Y direction (same Col)
	if !isSequentialInX {
		for i := 1; i < len(tiles); i++ {
			if tiles[i].Col != tiles[0].Col {
				isSequentialInX = false
				break
			}
		}
	} else {
		isSequentialInX = false // If they are sequential in X, they cannot be in Y
	}

	// If sequential in either direction, the adjacent elements must have consecutive indexes
	if isSequentialInX || isSequentialInY {
		sort.SliceStable(tiles, func(i, j int) bool {
			if isSequentialInX {
				return tiles[i].Col < tiles[j].Col
			}
			return tiles[i].Row < tiles[j].Row
		})

		// Check if adjacent elements have consecutive indexes
		for i := 1; i < len(tiles); i++ {
			if (isSequentialInX && tiles[i].Col-tiles[i-1].Col != 1) || (isSequentialInY && tiles[i].Row-tiles[i-1].Row != 1) {
				return false, false
			}
		}
	} else {
		return false, false
	}

	return true, false
}

// adjacentToPlacedTile checks if at least one new tile is adjacent to an already placed tile, unless the board is completely blank.
// It takes an activeGame of type Game and a newTiles of type MoveSlice as input.
// It returns a boolean value indicating whether at least one new tile is adjacent to an already placed tile.
func TestAdjacentToPlacedTile(activeGame Game, newTiles MoveSlice) bool {
	tempBoard := [15][15]string{}

	if reflect.DeepEqual(activeGame.Board, tempBoard) {
		return true
	}

	// Check if at least one new tile is adjacent to an already placed tile.
	for _, tile := range newTiles {
		row, col := tile.Col, tile.Row
		if row-1 >= 0 {
			if activeGame.Board[row-1][col] != "" && !containsTile(newTiles, row-1, col) {
				return true
			}
		}

		if row+1 <= 14 {
			if activeGame.Board[row+1][col] != "" && !containsTile(newTiles, row+1, col) {
				return true
			}
		}
		if col-1 >= 0 {
			if activeGame.Board[row][col-1] != "" && !containsTile(newTiles, row, col-1) {
				return true
			}
		}
		if col+1 <= 14 {
			if activeGame.Board[row][col+1] != "" && !containsTile(newTiles, row, col+1) {
				return true
			}
		}
	}
	return false
}

// containsTile checks if a tile with the given row and column exists in the given MoveSlice.
func containsTile(tiles MoveSlice, row, col int) bool {
	for _, tile := range tiles {
		if tile.Row == row && tile.Col == col {
			return true
		}
	}
	return false
}

// isFirstMove checks if this is the first move of the game.
// If it is, it checks that one tile submitted is on the center tile.
// It takes an activeGame of type Game and a newTiles of type MoveSlice as input.
// It returns an error if the move is invalid.
func isFirstMove(activeGame Game, newTiles MoveSlice) error {
	if activeGame.Board[7][7] == "" {
		return errors.New("first move must be on the center tile")
	}
	return nil
}
