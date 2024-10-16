package test

import (
	"github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/server/models"
	test "github.com/JohnGomes/Go-Test-Game-Web-Server/pkg/testing"
	"io"
	"testing"
)

func Test_Tape_Write(t *testing.T) {
	file, clean := test.CreateTempFile(t, "12345")
	defer clean()

	tape := &models.Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
