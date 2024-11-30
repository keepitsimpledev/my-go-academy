package main

import (
	"fmt"
	"go_academy/assignments/10-register"
	"strings"
	"time"
)

//nolint:gomnd
func main() {
	students := []register.Student{
		{
			Name:     "Kyle Alexander Kuzma",
			Birthday: toTimeDate(1995, 7, 24),
		},
		{
			Name:     "Jordan Anthony Poole",
			Birthday: toTimeDate(1999, 6, 19),
		},
		{
			Name:     "Tyus Robert Jones",
			Birthday: toTimeDate(1996, 5, 10),
		},
		{
			Name:     "Deni Avdija",
			Birthday: toTimeDate(2001, 1, 3),
		},
		{
			Name:     "Bilal Coulibaly",
			Birthday: toTimeDate(2004, 7, 26),
		},
		{
			Name:     "Corey James Kispert",
			Birthday: toTimeDate(1999, 3, 3),
		},
		{
			Name:     "Marvin Bagley III",
			Birthday: toTimeDate(1999, 3, 14),
		},
		{
			Name:     "Isaiah Maurice Livers",
			Birthday: toTimeDate(1998, 7, 28),
		},
		{
			Name:     "Landry Michael Shamet",
			Birthday: toTimeDate(1997, 3, 3),
		},
		{
			Name:     "Justin John Champagnie",
			Birthday: toTimeDate(2001, 6, 29),
		},
	}

	today := time.Now()

	var formattedStudentsInfo []string

	for _, student := range students {
		formattedStudentsInfo = append(formattedStudentsInfo, student.GetInfo(today))
	}

	fmt.Println(strings.Join(formattedStudentsInfo, "\n\n"))
}

func toTimeDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
