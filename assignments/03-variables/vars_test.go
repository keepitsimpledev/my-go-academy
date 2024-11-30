package main

import (
	"strings"
	"testing"
)

func TestPromptAndCheck10(t *testing.T) {
	numberCheckTests := []struct {
		description string
		input       string
		want        string
	}{
		{"0 is not in 1-10", "0", numberIsNotBetween},
		{"1 is in 1-10", "1", numberIsBetween},
		{"2 is in 1-10", "2", numberIsBetween},
		{"9 is in 1-10", "9", numberIsBetween},
		{"10 is in 1-10", "10", numberIsBetween},
		{"11 is in 1-10", "11", numberIsNotBetween},
	}

	for _, testCase := range numberCheckTests {
		t.Run(testCase.description, func(t *testing.T) {
			inputReader := strings.NewReader(testCase.input)

			got := PromptAndCheck10(inputReader)
			if got != testCase.want {
				t.Errorf("got: '%s'. want: '%s'. input: '%s'", got, testCase.want, testCase.input)
			}
		})
	}
}
