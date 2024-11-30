package poker

import (
	"fmt"
	"io"
	"time"
)

const blindFactor = 5

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}

// TexasHoldem manages a game of poker.
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewTexasHoldem returns a new game.
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

// Start will schedule blind alerts dependant on the number of players.
func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(blindFactor+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends the game, recording the winner.
func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
