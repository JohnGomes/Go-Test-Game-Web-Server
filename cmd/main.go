package main

import (
	"log"
	"net/http"

	d "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db"
	s "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
)

func main() {
	server := s.NewPlayerServer(d.NewInMemoryPlayerStore()) //&s.PlayerServer{Store: d.NewInMemoryPlayerStore(), Router: http.NewServeMux()}

	log.Fatal(http.ListenAndServe(":5000", server))
}
