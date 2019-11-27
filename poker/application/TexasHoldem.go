package application

import (
	"github.com/reuben-baek/learn-go-with-tests/poker/domain"
	"io"
	"time"
)

type TexasHoldem struct {
	alerter domain.BlindAlerter
	store   domain.PlayerStore
}

func NewTexasHoldem(alerter domain.BlindAlerter, store domain.PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter: alerter, store: store}
}

func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Second

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
