package main

import "testing"

func TestCombineStrings(t *testing.T) {
	want := "multiple strings variable"
	got := CombineStrings()

	if got != want {
		t.Errorf("want: %s. got %s.", want, got)
	}
}
