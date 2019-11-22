package main

import (
	"github.com/reuben-baek/learn-go-with-tests/player-server/endpoint"
	"github.com/reuben-baek/learn-go-with-tests/player-server/infrastructure"
)

var inMemoryPlayerStore = infrastructure.NewInMemoryPlayerStore()
var playerServer = endpoint.NewPlayerServer(inMemoryPlayerStore)
