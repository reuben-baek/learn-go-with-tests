package main

import (
	"github.com/reuben-baek/learn-go-with-tests/application/infrastructure/InMemoryPlayerStore"
	"github.com/reuben-baek/learn-go-with-tests/application/server"
	"log"
	"net/http"
)

var inMemoryPlayerStore = InMemoryPlayerStore.New()

func main() {
	server := &server.PlayerServer{Store: inMemoryPlayerStore}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("cound not listen on port 5000 %v", err)
	}
}
