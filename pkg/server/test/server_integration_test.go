package test

import (
	"encoding/json"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	helpers "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db/file_storage"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
)

func Test_RecordingWinsAndRetrievingFromFileStore(t *testing.T) {
	database, cleanup := helpers.CreateTempFile(t, `[]`)
	defer cleanup()
	store := file_storage.FilePlayerStore{Database: json.NewEncoder(database)}
	server := server.NewPlayerServer(&store)
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

func Test_RecordingWinsAndRetrievingFromInMemoryStore(t *testing.T) {
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
