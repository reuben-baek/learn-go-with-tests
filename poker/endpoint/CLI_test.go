package endpoint

import (
	"bytes"
	"strings"
	"testing"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &GameSpy{}
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		want := "Chris"
		if game.FinishedWith != want {
			t.Errorf("got %q, want %q", game.FinishedWith, want)
		}
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &GameSpy{}
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		want := "Cleo"
		if game.FinishedWith != want {
			t.Errorf("got %q, want %q", game.FinishedWith, want)
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}
		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, PlayerPrompt)

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, PlayerPrompt, BadPlayerInputErrMsg)
	})

}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
