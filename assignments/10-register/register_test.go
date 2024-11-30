package register_test

import (
	register "go_academy/assignments/10-register"
	"testing"
	"time"
)

func TestGetInfo(t *testing.T) {
	student := register.Student{
		Name:     "George Washington Carver",
		Birthday: time.Date(1864, time.Month(7), 12, 0, 0, 0, 0, time.UTC),
	}
	want := "full name: George Washington Carver\nbirthday: July 12, 1864\nage: 78"

	today := time.Date(1943, time.Month(1), 5, 0, 0, 0, 0, time.UTC)
	got := student.GetInfo(today)

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s\n\n", got, want)
	}
}
