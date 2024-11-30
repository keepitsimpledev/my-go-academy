package main

import (
	"fmt"
	"sort"
)

func InitArray() [10]int {
	return [10]int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}
}

func KeepOdd(numbers [10]int) []int {
	var out []int

	for _, number := range numbers {
		if number%2 == 1 {
			out = append(out, number)
		}
	}

	return out
}

func KeepEven(numbers [10]int) []int {
	var out []int

	for _, number := range numbers {
		if number%2 == 0 {
			out = append(out, number)
		}
	}

	return out
}

func SortAsc(numbers [10]int) [10]int {
	sort.Ints(numbers[:])
	return numbers
}

func SortDesc(numbers [10]int) [10]int {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers[:])))
	return numbers
}

func main() {
	arr := InitArray()

	ascending := SortAsc(arr)
	fmt.Println(ascending)

	descending := SortDesc(arr)
	fmt.Println(descending)

	ascOdd := KeepOdd(ascending)
	fmt.Println(ascOdd)

	ascEven := KeepEven(ascending)
	fmt.Println(ascEven)

	descOdd := KeepOdd(descending)
	fmt.Println(descOdd)

	descEven := KeepEven(descending)
	fmt.Println(descEven)
}
