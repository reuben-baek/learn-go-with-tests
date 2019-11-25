package endpoint

import "github.com/reuben-baek/learn-go-with-tests/poker/domain"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   domain.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() domain.League {
	return s.league
}
