package test

import (
	"encoding/json"
	"fmt"
	helpers "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing/stubs"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
)

func Test_GetLeagues(t *testing.T) {
	wantedLeague := []models.Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	_store := stubs.StubPlayerStore{nil, nil, wantedLeague}
	_server := server.NewPlayerServer(&_store)
	t.Run("Can Get Leagues", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		_server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, server.JSONContentType)

	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, contentType string) {
	if response.Result().Header.Get("content-type") != contentType {
		t.Errorf("response did not have content-type of application/json")
	}
}

func assertLeague(t *testing.T, got []models.Player, wantedLeague []models.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func getLeagueFromResponse(t *testing.T, response *httptest.ResponseRecorder) (league []models.Player) {
	t.Helper()
	err := json.NewDecoder(response.Body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
	}

	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/league", nil)
	return req
}

func Test_GetPlayers(t *testing.T) {

	_store := stubs.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		WinCalls: []string{},
	}
	_server := server.NewPlayerServer(&_store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		_server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		_server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		_server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, response.Code, http.StatusNotFound)

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}

	})

}

func TestStoreWins(t *testing.T) {
	_store := stubs.StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: []string{},
	}

	_server := server.NewPlayerServer(&_store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		_server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		helpers.AssertPlayerWin(t, &_store, "Chris")

	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
