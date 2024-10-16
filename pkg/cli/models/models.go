package models

import (
	"bufio"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	"io"
	"strings"
)

type CLI struct {
	Store models.PlayerStore
	In    *bufio.Scanner
}

func NewCLI(store models.PlayerStore, in io.Reader) *CLI {
	return &CLI{Store: store, In: bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	input := cli.readLine()
	cli.Store.RecordWin(extractWinner(input))
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
