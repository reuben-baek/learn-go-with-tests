package endpoint

import (
	"bufio"
	"github.com/reuben-baek/learn-go-with-tests/poker/domain"
	"io"
	"strings"
)

type CLI struct {
	playStore domain.PlayerStore
	in        *bufio.Scanner
}

func NewCLI(playStore domain.PlayerStore, in io.Reader) *CLI {
	return &CLI{playStore, bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()

	cli.playStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
