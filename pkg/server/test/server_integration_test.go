package test

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
)

func Test_RecordingWinsAndRetrievingThem(t *testing.T) {
	store := db.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response)
		want := []models.Player{{"Pepper", 3}}
		assertLeague(t, got, want)

	})

}
