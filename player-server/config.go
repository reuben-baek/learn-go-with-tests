package main

import (
	"github.com/reuben-baek/learn-go-with-tests/player-server/domain"
	"github.com/reuben-baek/learn-go-with-tests/player-server/endpoint"
	"github.com/reuben-baek/learn-go-with-tests/player-server/infrastructure"
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

var playerServer = endpoint.NewPlayerServer(fileSystemPlayerStore)
