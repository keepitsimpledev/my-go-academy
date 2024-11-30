package main

import (
	"testing"
	"testing/quick"
)

func TestRollDie(t *testing.T) {
	assertion := func() bool {
		result := RollDie()
		return result >= 1 && result <= 6
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount:      1000,
		MaxCountScale: 0,
		Rand:          nil,
		Values:        nil,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestRollDice(t *testing.T) {
	testRolls := []struct {
		rollA int
		rollB int
		want  string
	}{
		{6, 4, "6 4 - NEUTRAL"},
		{6, 3, "6 3 - NEUTRAL"},
		{3, 2, "3 2 - NEUTRAL"},
		{2, 2, "2 2 - NEUTRAL"},
		{5, 2, "5 2 - NATURAL"},
		{5, 6, "5 6 - NATURAL"},
		{1, 1, "1 1 - SNAKE-EYES-CRAPS"},
		{2, 1, "2 1 - LOSS-CRAPS"},
		{6, 6, "6 6 - LOSS-CRAPS"},
	}

	for _, testRoll := range testRolls {
		got := RollDice(testRoll.rollA, testRoll.rollB)
		if got != testRoll.want {
			t.Errorf("got: '%s'. want: '%s'. rolls: %d %d", got, testRoll.want, testRoll.rollA, testRoll.rollB)
		}
	}
}
