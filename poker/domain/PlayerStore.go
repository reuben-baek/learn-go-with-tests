package domain

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}
