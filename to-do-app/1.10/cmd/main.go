package main

import (
	"go_academy/to-do-app/1.10"
	"os"
)

func main() {
	print10.PrintThings(os.Stdout,
		"wash clothes", "fold clothes", "iron shirts", "wash dishes", "buy groceries",
		"mop the floor", "mow the lawn", "clean the bathtub", "pay bills", "take a nap")
}
