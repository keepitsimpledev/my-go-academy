package main

import (
	"strings"
	"testing"
)

func TestPromptAndFormatName(t *testing.T) {
	// arrange
	input := "Isa\nDalawa\nTatlo\n"
	rdr := strings.NewReader(input)

	// act
	got := PromptAndFormatName(rdr)

	// assert
	want := "Isa Dalawa Tatlo"
	if got != want {
		t.Errorf("got: '%s'. want '%s'.", got, want)
	}
}
