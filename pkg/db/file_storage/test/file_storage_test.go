package test

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db/file_storage"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	helpers "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing"
	"testing"
)

func Test_FileStorage(t *testing.T) {
	t.Run("Get League", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, `[{"Name": "Cleo", "Wins":10}, {"Name":"Chris", "Wins":33}]`)
		defer cleanup()

		store, _ := file_storage.NewFilePlayerStore(db)

		got := store.GetLeague()
		want := []models.Player{{"Cleo", 10}, {"Chris", 33}}
		got = store.GetLeague()

		helpers.AssertLeagueData(t, got, want)
	})

	t.Run("Get League ordered by score", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, `[{"Name": "Cranny", "Wins":32}, {"Name": "Jake", "Wins":1}, {"Name": "Cleo", "Wins":10}, {"Name":"Chris", "Wins":33}]`)
		defer cleanup()

		store, _ := file_storage.NewFilePlayerStore(db)

		want := []models.Player{{"Chris", 33}, {"Cranny", 32}, {"Cleo", 10}, {"Jake", 1}}
		got := store.GetLeague()

		helpers.AssertLeagueData(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, `[		{"Name": "Cleo", "Wins": 10},		{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()
		store, _ := file_storage.NewFilePlayerStore(db)

		got := store.GetPlayerScore("Chris")
		want := 33
		helpers.AssertEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, _ := file_storage.NewFilePlayerStore(db)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		helpers.AssertEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, _ := file_storage.NewFilePlayerStore(db)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		helpers.AssertEqual(t, got, want)

	})

	t.Run("works with an empty file", func(t *testing.T) {
		db, cleanup := helpers.CreateTempFile(t, "")
		defer cleanup()

		_, err := file_storage.NewFilePlayerStore(db)

		if err != nil {
			t.Errorf("Failed to make NewFilePlayerStore with an empry file : %s", err)
		}
	})
}
