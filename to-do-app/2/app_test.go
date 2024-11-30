package app_test

import (
	"bufio"
	"bytes"
	"fmt"
	app "go_academy/to-do-app/2"
	"strings"
	"testing"
)

const inputPromptTemplate = "Enter a number to perform its correspoding action:\n" +
	"[1] Add a new To-Do list item\n" +
	"[2] View a To-Do list item\n" +
	"[3] Update a To-Do list item\n" +
	"[4] Delete a To-Do list item\n" +
	"[5] Exit the To-Do list application\n\n" +
	"input %s received\n\n"
const enterTheTask = "Enter the task:\n\n"
const exiting = "exiting\n"

func TestRun(t *testing.T) {
	outputBuffer := bytes.Buffer{}
	inputReader := strings.NewReader("5\n")

	app.Run(inputReader, &outputBuffer)

	got := outputBuffer.String()

	want := fmt.Sprintf(inputPromptTemplate+exiting, "5")

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}

func TestPromptForAction(t *testing.T) {
	promptTest := func(tb testing.TB, input string, wantAction int) {
		tb.Helper()

		// arrange
		outputBuffer := bytes.Buffer{}
		wantOutput := fmt.Sprintf(inputPromptTemplate, input)

		inputString := fmt.Sprintf("%s\n", input)
		inputReader := bufio.NewReader(strings.NewReader(inputString))

		// act
		gotAction := app.PromptForAction(inputReader, &outputBuffer)
		gotOutput := outputBuffer.String()

		// assert
		if gotOutput != wantOutput {
			t.Errorf("\ngot: %s\nwant: %s", gotOutput, wantOutput)
		}

		if gotAction != wantAction {
			t.Errorf("\ngot: %s\nwant: %s", gotOutput, wantOutput)
		}
	}

	tests := []struct {
		description, input string
		wantAction         int
	}{
		{"invalid input of '-1'", "-1", -1},
		{"invalid input of '0'", "0", -1},
		{"valid input of '1'", "1", 1},
		{"valid input of '2'", "2", 2},
		{"valid input of '3'", "3", 3},
		{"valid input of '4'", "4", 4},
		{"invalid input of '5'", "5", 5},
		{"invalid input of '6'", "6", -1},
		{"invalid input of '6'", "6", -1},
		{"invalid input of 'word'", "word", -1},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			promptTest(t, test.input, test.wantAction)
		})
	}
}

func TestAddAndRead(t *testing.T) {
	inputString := "1\ndo laundry\n2\n5\n"
	inputReader := strings.NewReader(inputString)

	buffer := bytes.Buffer{}

	app.Run(inputReader, &buffer)

	got := buffer.String()

	want := fmt.Sprintf(inputPromptTemplate, "1")
	want += "Enter the task:\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "2")
	want += "To-Do list:\n1. do laundry - Not Started\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "5")
	want += exiting

	assert(t, got, want)
}

func TestAddUpdateAndRead(t *testing.T) {
	inputString := "1\nlaundry\n3\n1\ndishes\nin-progress\n2\n5\n"
	inputReader := strings.NewReader(inputString)

	buffer := bytes.Buffer{}

	app.Run(inputReader, &buffer)

	got := buffer.String()

	want := fmt.Sprintf(inputPromptTemplate, "1")
	want += enterTheTask
	want += fmt.Sprintf(inputPromptTemplate, "3")
	want += "Enter the number to update:\n"
	want += "Enter the item update:\n"
	want += "Enter the status update:\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "2")
	want += "To-Do list:\n1. dishes - in-progress\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "5")
	want += exiting

	assert(t, got, want)
}

func TestUpdateNotFound(t *testing.T) {
	inputString := "3\n1\ndishes\nin-progress\n5\n"
	inputReader := strings.NewReader(inputString)

	buffer := bytes.Buffer{}

	app.Run(inputReader, &buffer)

	got := buffer.String()

	want := fmt.Sprintf(inputPromptTemplate, "3")
	want += "Enter the number to update:\n"
	want += "Enter the item update:\n"
	want += "Enter the status update:\n\n"
	want += "Not found.\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "5")
	want += exiting

	assert(t, got, want)
}

func TestAddDeleteAndRead(t *testing.T) {
	inputString := "1\nlaundry\n4\n1\n2\n5\n"
	inputReader := strings.NewReader(inputString)

	buffer := bytes.Buffer{}

	app.Run(inputReader, &buffer)

	got := buffer.String()

	want := fmt.Sprintf(inputPromptTemplate, "1")
	want += enterTheTask
	want += fmt.Sprintf(inputPromptTemplate, "4")
	want += "Enter the number to delete:\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "2")
	want += "To-Do list is empty\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "5")
	want += exiting

	assert(t, got, want)
}

func TestDeleteNotFound(t *testing.T) {
	inputString := "4\n1\n5\n"
	inputReader := strings.NewReader(inputString)

	buffer := bytes.Buffer{}

	app.Run(inputReader, &buffer)

	got := buffer.String()

	want := fmt.Sprintf(inputPromptTemplate, "4")
	want += "Enter the number to delete:\n\n"
	want += "Not found.\n\n"
	want += fmt.Sprintf(inputPromptTemplate, "5")
	want += exiting

	assert(t, got, want)
}

func assert(tb testing.TB, got, want any) {
	tb.Helper()

	if got != want {
		tb.Errorf("got:\n%v\n\nwant:\n%v", got, want)
	}
}
