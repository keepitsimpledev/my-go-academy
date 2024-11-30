package main

import (
	json11 "go_academy/to-do-app/1.11"
	"os"
)

//nolint:gomnd
func main() {
	first := json11.Task{Number: 1, Name: "wash clothes"}
	second := json11.Task{Number: 2, Name: "fold clothes"}
	third := json11.Task{Number: 3, Name: "iron shirts"}
	fourth := json11.Task{Number: 4, Name: "wash dishes"}
	fifth := json11.Task{Number: 5, Name: "buy groceries"}
	sixth := json11.Task{Number: 6, Name: "mop the floor"}
	seventh := json11.Task{Number: 7, Name: "mow the lawn"}
	eighth := json11.Task{Number: 8, Name: "clean the bathtub"}
	ninth := json11.Task{Number: 9, Name: "pay bills"}
	tenth := json11.Task{Number: 10, Name: "take a nap"}

	json11.PrintTasksToJSON(os.Stdout,
		first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth)
}
