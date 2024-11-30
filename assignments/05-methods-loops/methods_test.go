package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestPromptForNumbers(t *testing.T) {
	// arrange
	var input string

	numbers := [9]int{2, 12, 22, 7, 7}
	for _, number := range numbers {
		addToInput(&input, number)
	}

	reader := strings.NewReader(input)

	// act
	got := PromptForNumbers(reader, 1, 9)

	// assert
	want := [3]int{2, 7, 7}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: '%d'. want '%d'.", got, want)
	}
}

func TestSumNineNumbers(t *testing.T) {
	// arrange
	var input string

	numbers := []int{2, 12, 22, 7, 7, 8, 100, 20, 25, 7, 101, 99, 98, 1000, 100, 200, 97, 9999, 999}
	for _, number := range numbers {
		addToInput(&input, number)
	}

	reader := strings.NewReader(input)

	// act
	got := SumNineNumbers(reader)

	// assert
	want := 1459
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: '%d'. want '%d'.", got, want)
	}
}

func addToInput(inputBuffer *string, newInput int) {
	*inputBuffer += fmt.Sprintf("%d\n", newInput)
}
