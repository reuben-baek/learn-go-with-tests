package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/reuben-baek/learn-go-with-tests/player-server/domain"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   domain.League
}

func (f *FileSystemPlayerStore) GetLeague() domain.League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		newPlayer := domain.Player{Name: name, Wins: 1}
		f.league = f.league.Add(newPlayer)
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}

func NewLeague(rdr io.Reader) ([]domain.Player, error) {
	var league []domain.Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) domain.PlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}
