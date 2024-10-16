package db

import "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetLeague() models.League {
	var league models.League
	for name, wins := range i.store {
		league = append(league, models.Player{name, wins})
	}
	return league
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
