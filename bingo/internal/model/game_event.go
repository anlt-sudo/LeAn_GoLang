package model

type GameEvent struct {
	EventType string
	Message   string
	Board     [][]string
	BingoPos  [][2]int
}