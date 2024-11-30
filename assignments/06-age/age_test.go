package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPromptForYear(t *testing.T) {
	input := ""
	addToInput(&input, -1)
	addToInput(&input, 1975)

	reader := strings.NewReader(input)
	got := PromptForYear(reader)
	want := 1975

	if got != want {
		t.Errorf("got: '%d'. want: '%d'.", got, want)
	}
}

func TestPromptForMonth(t *testing.T) {
	input := ""
	addToInput(&input, -1)
	addToInput(&input, 0)
	addToInput(&input, 13)
	addToInput(&input, 14)
	addToInput(&input, 2)

	reader := strings.NewReader(input)
	got := PromptForMonth(reader)
	want := 2

	if got != want {
		t.Errorf("got: '%d'. want: '%d'.", got, want)
	}
}

func TestPromptForDay(t *testing.T) {
	input := ""
	addToInput(&input, 15)

	reader := strings.NewReader(input)
	got := PromptForDay(reader)
	want := 15

	if got != want {
		t.Errorf("got: '%d'. want: '%d'.", got, want)
	}
}

func TestCalculateAge(t *testing.T) {
	// arrange
	input := ""
	addToInput(&input, 1994)
	addToInput(&input, 1)
	addToInput(&input, 19)
	reader := strings.NewReader(input)
	now := time.Date(2024, 3, 30, 0, 0, 0, 0, time.UTC)

	// act
	got := CalculateAge(reader, now)

	// assert
	want := 30
	if got != want {
		t.Errorf("got: '%d'. want: '%d'.", got, want)
	}
}

func addToInput(inputBuffer *string, newInput int) {
	*inputBuffer += fmt.Sprintf("%d\n", newInput)
}
