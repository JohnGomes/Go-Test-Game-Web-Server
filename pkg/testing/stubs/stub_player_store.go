package stubs

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   models.League
}

func (s *StubPlayerStore) GetLeague() models.League {
	return s.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}
