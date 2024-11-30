package main

import (
	"fmt"
	"github.com/bearbin/go-age"
	"io"
	"os"
	"time"
)

func PromptForYear(in io.Reader) int {
	if in == nil {
		in = os.Stdin
	}

	year := -1

	for year < 0 {
		fmt.Printf("Enter birth year: ")
		fmt.Fscanln(in, &year)

		if year < 0 {
			fmt.Println("Invalid year value (must be 0 or greater)")
		}
	}

	return year
}

func PromptForMonth(in io.Reader) int {
	if in == nil {
		in = os.Stdin
	}

	var month int

	for month < 1 || month > 12 {
		fmt.Printf("Enter birth month (1-12): ")
		fmt.Fscanln(in, &month)

		if month < 1 || month > 12 {
			fmt.Println("Invalid month value")
		}
	}

	return month
}

func PromptForDay(in io.Reader) int {
	if in == nil {
		in = os.Stdin
	}

	var day int

	fmt.Printf("Enter birth day: ")
	fmt.Fscanln(in, &day)

	return day
}

func CalculateAge(in io.Reader, now time.Time) int {
	birthyear := PromptForYear(in)
	birthmonth := PromptForMonth(in)
	birthday := PromptForDay(in)
	birthdate := time.Date(birthyear, time.Month(birthmonth), birthday, 0, 0, 0, 0, time.UTC)

	return age.AgeAt(birthdate, now)
}

func main() {
	age := CalculateAge(os.Stdin, time.Now())
	fmt.Printf("Age: %d\n", age)
}
