package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/reuben-baek/learn-go-with-tests/player-server/domain"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() domain.League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()

	player := league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		newPlayer := domain.Player{Name: name, Wins: 1}
		league = league.Add(newPlayer)
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
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
	return &FileSystemPlayerStore{database}
}
