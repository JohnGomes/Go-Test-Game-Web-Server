package file_storage

import (
	"encoding/json"
	"fmt"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	"io"
	"log"
	"os"
	"sort"
)

type FilePlayerStore struct {
	Database *json.Encoder
	League   models.League
}

func NewFilePlayerStore(file *os.File) (*FilePlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("Error initializing player DB: %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FilePlayerStore{
		Database: json.NewEncoder(&models.Tape{file}),
		League:   league}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file infor from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func (store *FilePlayerStore) GetLeague() models.League {
	sort.Slice(store.League, func(i, j int) bool {
		return store.League[i].Wins > store.League[j].Wins
	})
	return store.League
}

func (store *FilePlayerStore) GetPlayerScore(name string) int {
	player := store.League.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (store *FilePlayerStore) RecordWin(name string) {
	player := store.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		store.League = append(store.League, models.Player{name, 1})
	}

	//store.Database.Seek(0, io.SeekStart)
	store.Database.Encode(store.League)
}

func NewLeague(rdr io.Reader) ([]models.Player, error) {
	var league []models.Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
		log.Default().Println(err)
	}

	return league, err
}
