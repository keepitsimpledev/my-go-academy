package main

import (
	"fmt"
	"io"
	"os"
)

const numberIsBetween = "the number lies between 1 and 10"
const numberIsNotBetween = "the number is not between 1 and 10"

func PromptAndCheck10(in io.Reader) string {
	if in == nil {
		in = os.Stdin
	}

	var number int

	fmt.Print("Enter a number: ")
	fmt.Fscanln(in, &number)

	if number >= 1 && number <= 10 {
		return numberIsBetween
	}

	return numberIsNotBetween
}

func main() {
	fmt.Println(PromptAndCheck10(os.Stdin))
}
