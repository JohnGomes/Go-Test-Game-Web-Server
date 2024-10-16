package test

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/cli/models"
	helpers "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing"
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing/stubs"
	"strings"
	"testing"
)

func Test_CLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &stubs.StubPlayerStore{}

		cli := models.NewCLI(store, in)
		cli.PlayPoker()

		helpers.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &stubs.StubPlayerStore{}

		cli := models.NewCLI(store, in)
		cli.PlayPoker()

		helpers.AssertPlayerWin(t, store, "Cleo")
	})
}
