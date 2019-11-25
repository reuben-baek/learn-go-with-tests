package endpoint

import (
	"bufio"
	"fmt"
	"github.com/reuben-baek/learn-go-with-tests/poker/domain"
	"io"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	playStore domain.PlayerStore
	in        *bufio.Scanner
	out       io.Writer
	alerter   BlindAlerter
}

func NewCLI(playStore domain.PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{playStore, bufio.NewScanner(in), out, alerter}
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(cli.readLine())
	cli.scheduleBlindAlerts(numberOfPlayers)
	userInput := cli.readLine()
	cli.playStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
