package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	d "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db"
)

func Test_RecordingWinsAndRetrievingThem(t *testing.T) {
	store := d.NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")

}
