package main

import (
	"fmt"
	"io"
	"os"
)

const (
	numbersToPromptFor = 3
	singleDigitLo      = 1
	singleDigitHi      = 9
	doubleDigitLo      = 10
	doubleDigitHi      = 99
	tripleDigitLo      = 100
	tripleDigitHi      = 999
)

func PromptForNumbers(in io.Reader, low, high int) [numbersToPromptFor]int {
	if in == nil {
		in = os.Stdin
	}

	var numbers [numbersToPromptFor]int

	var received int

	for i := 0; i < numbersToPromptFor; {
		fmt.Printf("Enter a number between %d and %d: ", low, high)
		fmt.Fscanln(in, &received)

		if received >= low && received <= high {
			numbers[i] = received
			i++
		} else {
			fmt.Printf("%d is not between %d and %d\n", received, low, high)
		}
	}

	return numbers
}

func SumNineNumbers(in io.Reader) int {
	singles := PromptForNumbers(in, singleDigitLo, singleDigitHi)
	doubles := PromptForNumbers(in, doubleDigitLo, doubleDigitHi)
	triples := PromptForNumbers(in, tripleDigitLo, tripleDigitHi)

	sum := 0
	for _, num := range singles {
		sum += num
	}

	for _, num := range doubles {
		sum += num
	}

	for _, num := range triples {
		sum += num
	}

	return sum
}

func main() {
	fmt.Printf("the sum is: %v\n", SumNineNumbers(os.Stdin))
}
