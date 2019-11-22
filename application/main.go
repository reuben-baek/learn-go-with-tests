package main

import (
	"github.com/reuben-baek/learn-go-with-tests/application/server"
	"log"
	"net/http"
)

var inMemoryPlayerStore = &InMemoryPlayerStore{
	scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	},
}

func main() {
	server := &server.PlayerServer{Store: inMemoryPlayerStore}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("cound not listen on port 5000 %v", err)
	}
}

type InMemoryPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
