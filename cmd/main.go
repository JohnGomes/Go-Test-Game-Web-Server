package main

import (
	"log"
	"net/http"

	d "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db"
	s "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server"
)

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	server := &s.PlayerServer{d.NewInMemoryPlayerStore()}
	// server.ServeHttp()
	log.Fatal(http.ListenAndServe(":5000", server))
}
