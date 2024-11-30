package main

import (
	"fmt"
	"io"
	"os"
)

func PromptAndFormatName(in io.Reader) string {
	if in == nil {
		in = os.Stdin
	}

	var first string

	var middle string

	var last string

	fmt.Print("Enter first name: ")
	fmt.Fscanln(in, &first)

	fmt.Print("Enter middle name: ")
	fmt.Fscanln(in, &middle)

	fmt.Print("Enter last name: ")
	fmt.Fscanln(in, &last)

	return fmt.Sprintf("%s %s %s", first, middle, last)
}

func main() {
	fmt.Println(PromptAndFormatName(os.Stdin))
}
