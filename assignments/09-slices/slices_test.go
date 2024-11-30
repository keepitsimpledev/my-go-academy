package main

import (
	"strings"
	"testing"
)

func TestPromptForName(t *testing.T) {
	// arrange
	input := "Henry\nDavid\nThoreau\n"
	rdr := strings.NewReader(input)

	// act
	got := PromptForName(rdr)

	// assert
	want := Fullname{first: "Henry", middle: "David", last: "Thoreau"}
	if got != want {
		t.Errorf("got:\n'%s'\nwant:\n'%s'", got, want)
	}
}

func TestFormatName(t *testing.T) {
	// arrange
	name := Fullname{first: "Henry", middle: "David", last: "Thoreau"}

	// act
	got := FormatName(name)

	// assert
	want := `full-name : Henry David Thoreau
middle-name : David
surname : Thoreau`
	if got != want {
		t.Errorf("got:\n'%s'\nwant:\n'%s'", got, want)
	}
}
