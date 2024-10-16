package main

import (
	"fmt"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/cli/models"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/db/file_storage"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {

	store, cleanup, err := file_storage.FilePlayStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	models.NewCLI(store, os.Stdin).PlayPoker()
}
