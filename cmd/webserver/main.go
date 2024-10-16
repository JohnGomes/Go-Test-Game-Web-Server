package main

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db/file_storage"
	s "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	store, cleanup, err := file_storage.FilePlayStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	server := s.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
