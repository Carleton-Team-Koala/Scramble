package models

type Game struct {
	GameID           string                `json:"GameID"`
	Board            [15][15]string        `json:"Board"`
	AvailableLetters map[string]int        `json:"LetterDistribution"`
	Players          map[string]PlayerInfo `json:"Players"`
	CurrentPlayer    string                `json:"CurrentPlayer"`
	PlayerList       []string              `json:"-"`
	TotalMoves       int                   `json:"TotalMoves"`
	GameStarted      bool                  `json:"GameStarted"`
	GameOver         bool                  `json:"GameOver"`
	Winner           string                `json:"Winner"`
}

type PlayerInfo struct {
	Score int      `json:"score"`
	Hand  []string `json:"hand"`
}

type UpdateGameResp struct {
	PlayerName string `json:"playerName"`
	Updates    []Move `json:"updates"`
}

type WordScore struct {
	Valid bool
	Score int
}

type PlayerNameResp struct {
	PlayerName string `json:"playerName"`
}

type Move struct {
	Letter string `json:"letter"`
	Col    int    `json:"yLoc"`
	Row    int    `json:"xLoc"`
}

type SkipTurnResp struct {
	CurrentPlayer string `json:"CurrentPlayer"`
	Message       string `json:"message"`
}

type RefreshResp struct {
	CurrentPlayer string   `json:"CurrentPlayer"`
	NewHand       []string `json:"hand"`
}
