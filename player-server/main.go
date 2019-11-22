package main

import (
	"github.com/reuben-baek/learn-go-with-tests/player-server/endpoint"
	"github.com/reuben-baek/learn-go-with-tests/player-server/infrastructure"
	"log"
	"net/http"
)

var inMemoryPlayerStore = infrastructure.NewInMemoryPlayerStore()

func main() {
	server := endpoint.NewPlayerServer(inMemoryPlayerStore)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("cound not listen on port 5000 %v", err)
	}
}
