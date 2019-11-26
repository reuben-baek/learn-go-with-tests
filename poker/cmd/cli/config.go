package main

import (
	"github.com/reuben-baek/learn-go-with-tests/poker/application"
	"github.com/reuben-baek/learn-go-with-tests/poker/domain"
	"github.com/reuben-baek/learn-go-with-tests/poker/endpoint"
	"github.com/reuben-baek/learn-go-with-tests/poker/infrastructure"
	"log"
	"os"
)

const dbFileName = "game.db.json"

var fileSystemPlayerStore, fileSystemPlayerStoreClose = func() (domain.PlayerStore, func()) {
	store, close, err := infrastructure.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	return store, close
}()

var game = application.NewTexasHoldem(domain.BlindAlerterFunc(domain.StdOutAlerter), fileSystemPlayerStore)
var cli = endpoint.NewCLI(os.Stdin, os.Stdout, game)
