package application

import "github.com/reuben-baek/learn-go-with-tests/poker/domain"

type StubPlayerStore struct {
	scores   map[string]int
	WinCalls []string
	league   domain.League
}

func NewStubPlayerStore(scores map[string]int, winCalls []string, league domain.League) *StubPlayerStore {
	return &StubPlayerStore{scores: scores, WinCalls: winCalls, league: league}
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() domain.League {
	return s.league
}
