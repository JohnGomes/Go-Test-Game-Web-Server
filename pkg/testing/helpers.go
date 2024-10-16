package test

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing/stubs"
	"os"
	"reflect"
	"testing"
)

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func AssertEqual(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertLeagueData(t *testing.T, got []models.Player, want []models.Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertPlayerWin(t *testing.T, store *stubs.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}
