package InMemoryPlayerStore

import "github.com/reuben-baek/learn-go-with-tests/application/domain"

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []domain.Player {
	var league []domain.Player
	for name, wins := range i.scores {
		league = append(league, domain.Player{name, wins})
	}
	return league
}

func New() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}
