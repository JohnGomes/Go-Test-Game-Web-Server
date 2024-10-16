package main

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db/file_storage"
	"log"
	"net/http"
	"os"

	s "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Error opening database file: %s %v", dbFileName, err)
	}

	store, err := file_storage.NewFilePlayerStore(db) //&s.PlayerServer{Store: d.NewInMemoryPlayerStore(), Router: http.NewServeMux()}

	if err != nil {
		log.Fatalf("unexpected error %v", err)
	}

	server := s.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
