package register

import (
	"fmt"
	"github.com/bearbin/go-age"
	"time"
)

type Student struct {
	Name     string
	Birthday time.Time
}

func (s Student) GetInfo(today time.Time) string {
	month := s.Birthday.Month().String()
	formattedBirthday := fmt.Sprintf("%s %d, %d", month, s.Birthday.Day(), s.Birthday.Year())
	age := age.AgeAt(s.Birthday, today)

	return fmt.Sprintf("full name: %s\nbirthday: %s\nage: %d", s.Name, formattedBirthday, age)
}
