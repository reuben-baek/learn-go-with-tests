package main

import (
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

var fileSystemPlayerStore = infrastructure.NewFileSystemPlayerStore(db)
var playerServer = endpoint.NewPlayerServer(fileSystemPlayerStore)
