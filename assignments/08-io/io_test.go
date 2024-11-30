package main

import (
	"errors"
	"os"
	"slices"
	"testing"
)

func TestCreateAndClose(t *testing.T) {
	CreateFile()

	if !fileExists(t, Filename) {
		t.Errorf("expected to find %s, but did not find it", Filename)
	}

	DeleteFile()

	if fileExists(t, Filename) {
		t.Errorf("%s should have been deleted, but was not", Filename)
	}
}

func TestReadAndWrite(t *testing.T) {
	want := "TestReadAndWrite"
	file := CreateFile()
	WriteToFile(want, file)

	got := ReadFromFile()

	DeleteFile()

	if got != want {
		t.Errorf("got: '%s'. want: '%s'", got, want)
	}
}

func TestWriteAndReadFromFile(t *testing.T) {
	citiesList := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}
	want := []string{"Abu Dhabi", "Caracas", "Hanoi", "London", "Montevideo", "Vatican City", "Washington D.C."}
	got := WriteAndReadFromFile(citiesList)

	if !slices.Equal(got, want) {
		t.Errorf("got:\n%s\nwant:\n%s\n\n", got, want)
	}
}

func fileExists(tb testing.TB, path string) bool {
	tb.Helper()

	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	panic(err)
}
