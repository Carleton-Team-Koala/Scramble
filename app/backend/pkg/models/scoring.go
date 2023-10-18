package models

func scoring(activeGame Game, newTiles string){
	var word = ""
	for i:= 0, i< len(newTiles), i++{
		word = word + newTiles[i]


	}
}

func checkLeft(x int, y int, activeGame Game, newTiles string, wordSoFar string){
	if x > 0 {
		if activeGame.Board[x-1][y] != ""{
			checkLeft(x-1. y, activactiveGame, newTiles, activeGame.Board[x-1][y] + wordSoFar)
		}
	}
}