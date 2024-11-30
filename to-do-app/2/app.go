package app

import (
	"bufio"
	"fmt"
	"go_academy/to-do-app/2/datastructure"
	"io"
	"strconv"
	"strings"
)

const inputPrompt = `Enter a number to perform its correspoding action:
[1] Add a new To-Do list item
[2] View a To-Do list item
[3] Update a To-Do list item
[4] Delete a To-Do list item
[5] Exit the To-Do list application

`

const add = 1
const read = 2
const update = 3
const deleteItem = 4
const exit = 5

func Run(inputReader io.Reader, outputWriter io.Writer) {
	bufioInputReader := bufio.NewReader(inputReader)
	todoList := ds.NewList()

	action := 2
	for isValidAction(action) {
		action = PromptForAction(bufioInputReader, outputWriter)

		switch action {
		case add:
			CmdAdd(&todoList, outputWriter, bufioInputReader)
		case read:
			CmdRead(&todoList, outputWriter)
		case update:
			CmdUpdate(&todoList, outputWriter, bufioInputReader)
		case deleteItem:
			CmdDelete(&todoList, outputWriter, bufioInputReader)
		case exit:
			fmt.Fprintln(outputWriter, "exiting")
			return
		default:
			fmt.Fprintln(outputWriter, "action not found - exiting")
			return
		}

		fmt.Fprintln(outputWriter)
	}
}

func PromptForAction(inputReader *bufio.Reader, outputWriter io.Writer) int {
	fmt.Fprint(outputWriter, inputPrompt)

	receivedUserInput := readLine(inputReader)

	fmt.Fprintf(outputWriter, "input %s received\n\n", receivedUserInput)

	receivedInt, err := strconv.Atoi(receivedUserInput)
	if err != nil || !isValidAction(receivedInt) {
		return -1
	}

	return receivedInt
}

func isValidAction(action int) bool {
	return action >= add && action <= exit
}

func CmdAdd(todoList *ds.TodoList, outputWriter io.Writer, inputReader *bufio.Reader) {
	fmt.Fprintln(outputWriter, "Enter the task:")

	item := readLine(inputReader)

	todoList.Add(item, "Not Started")
}

func CmdRead(todoList *ds.TodoList, outputWriter io.Writer) {
	fmt.Fprint(outputWriter, todoList.GetAll())
}

func CmdUpdate(todoList *ds.TodoList, outputWriter io.Writer, inputReader *bufio.Reader) {
	fmt.Fprintln(outputWriter, "Enter the number to update:")

	number := readLine(inputReader)

	numberInt, atoiErr := strconv.Atoi(number)
	if atoiErr != nil {
		fmt.Fprintf(outputWriter, "invalid input: '%s'\n", number)
		return
	}

	fmt.Fprintln(outputWriter, "Enter the item update:")

	item := readLine(inputReader)

	fmt.Fprintln(outputWriter, "Enter the status update:")

	status := readLine(inputReader)

	updateErr := todoList.Update(numberInt-1, item, status)
	if updateErr != nil {
		fmt.Fprintf(outputWriter, "\nNot found.\n")
	}
}

func CmdDelete(todoList *ds.TodoList, outputWriter io.Writer, inputReader *bufio.Reader) {
	fmt.Fprintln(outputWriter, "Enter the number to delete:")

	number := readLine(inputReader)

	numberInt, atoiErr := strconv.Atoi(number)
	if atoiErr != nil {
		fmt.Fprintf(outputWriter, "invalid input: '%s'\n", number)
		return
	}

	updateErr := todoList.Delete(numberInt - 1)
	if updateErr != nil {
		fmt.Fprintf(outputWriter, "\nNot found.\n")
	}
}

func readLine(inputReader *bufio.Reader) string {
	item, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(item, "\n")
}
