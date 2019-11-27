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

var db = func() *os.File {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	return db
}()

var fileSystemPlayerStore = func() domain.PlayerStore {
	store, err := infrastructure.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	return store
}()

var game = application.NewTexasHoldem(domain.BlindAlerterFunc(domain.Alerter), fileSystemPlayerStore)

var playerServer = func() *endpoint.PlayerServer {
	server, err := endpoint.NewPlayerServer(fileSystemPlayerStore, game)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	return server
}()
