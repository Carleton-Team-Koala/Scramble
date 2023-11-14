package models

type Game struct {
	GameID           string                `json:"GameID"`
	Board            [15][15]string        `json:"Board"`
	AvailableLetters map[string]int        `json:"LetterDistribution"`
	Players          map[string]PlayerInfo `json:"Players"`
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

type MoveSlice []Move

func (m MoveSlice) Len() int {
	return len(m)
}

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
