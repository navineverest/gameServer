package main

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
	GetLeaguePlayers() []Player
}
