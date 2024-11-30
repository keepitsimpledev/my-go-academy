package main

import (
	"fmt"
	"math/rand/v2"
)

const dieMax = 6
const numRolls = 50

func RollDie() int {
	return rand.IntN(dieMax) + 1
}

//nolint:gomnd
func Result(a, b int) string {
	sum := a + b

	switch sum {
	case 7:
		fallthrough
	case 11:
		return "NATURAL"
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 3:
		fallthrough
	case 12:
		return "LOSS-CRAPS"
	default:
		return "NEUTRAL"
	}
}

func RollDice(rollA, rollB int) string {
	result := Result(rollA, rollB)

	return fmt.Sprintf("%d %d - %s", rollA, rollB, result)
}

func main() {
	for i := 0; i < numRolls; i++ {
		fmt.Println(RollDice(RollDie(), RollDie()))
	}
}
