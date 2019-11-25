package endpoint

import (
	"bufio"
	"github.com/reuben-baek/learn-go-with-tests/poker/domain"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playStore domain.PlayerStore
	in        *bufio.Scanner
	alerter   BlindAlerter
}

func NewCLI(playStore domain.PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{playStore, bufio.NewScanner(in), alerter}
}

func (cli *CLI) PlayPoker() {
	cli.scheduleBlindAlerts()
	userInput := cli.readLine()
	cli.playStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
