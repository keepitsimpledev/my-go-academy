package main

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	want := [10]int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}
	got := InitArray()

	assertArrays(t, want[:], got[:])
}
func TestSortAsc(t *testing.T) {
	want := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr := InitArray()
	got := SortAsc(arr)

	assertArrays(t, want[:], got[:])
}
func TestSortDesc(t *testing.T) {
	want := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	arr := InitArray()
	got := SortDesc(arr)

	assertArrays(t, want[:], got[:])
}
func TestKeepOdd(t *testing.T) {
	want := []int{1, 3, 5, 7, 9}

	arr := InitArray()
	got := KeepOdd(arr)

	assertArrays(t, want, got)
}
func TestKeepEven(t *testing.T) {
	want := []int{10, 2, 4, 6, 8}

	arr := InitArray()
	got := KeepEven(arr)

	assertArrays(t, want, got)
}

func assertArrays(tb testing.TB, want, got []int) {
	tb.Helper()

	if !reflect.DeepEqual(want, got) {
		tb.Errorf("got: %v. want: %v.", got, want)
	}
}
