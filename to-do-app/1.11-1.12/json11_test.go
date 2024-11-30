package json11and12_test

import (
	"bytes"
	json11 "go_academy/to-do-app/1.11-1.12"
	"os"
	"testing"
)

func TestPrintTasksToJSON(t *testing.T) {
	want := `[{"Number":1,"Name":"amp"},{"Number":2,"Name":"banana"},{"Number":3,"Name":"car"},` +
		`{"Number":4,"Name":"doorknob"},{"Number":5,"Name":"egg"},{"Number":6,"Name":"feather"},` +
		`{"Number":7,"Name":"gold"},{"Number":8,"Name":"hanger"},{"Number":9,"Name":"inchworm"},` +
		`{"Number":10,"Name":"jellybean"}]` + "\n"

	one := json11.Task{Number: 1, Name: "amp"}
	two := json11.Task{Number: 2, Name: "banana"}
	three := json11.Task{Number: 3, Name: "car"}
	four := json11.Task{Number: 4, Name: "doorknob"}
	five := json11.Task{Number: 5, Name: "egg"}
	six := json11.Task{Number: 6, Name: "feather"}
	seven := json11.Task{Number: 7, Name: "gold"}
	eight := json11.Task{Number: 8, Name: "hanger"}
	nine := json11.Task{Number: 9, Name: "inchworm"}
	ten := json11.Task{Number: 10, Name: "jellybean"}

	buffer := bytes.Buffer{}
	json11.PrintTasksToJSON(&buffer, one, two, three, four, five, six, seven, eight, nine, ten)

	got := buffer.String()

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}

func ExamplePrintTasksToJSON() {
	taskA := json11.Task{Number: 1, Name: "do laundry"}
	taskB := json11.Task{Number: 2, Name: "wash dishes"}
	taskC := json11.Task{Number: 3, Name: "clean car"}
	json11.PrintTasksToJSON(os.Stdout, taskA, taskB, taskC)
	// Output:
	// [{"Number":1,"Name":"do laundry"},{"Number":2,"Name":"wash dishes"},{"Number":3,"Name":"clean car"}]
}

func BenchmarkPrintThings(b *testing.B) {
	one := json11.Task{Number: 1, Name: "amp"}
	two := json11.Task{Number: 2, Name: "banana"}
	three := json11.Task{Number: 3, Name: "car"}
	four := json11.Task{Number: 4, Name: "doorknob"}
	five := json11.Task{Number: 5, Name: "egg"}
	six := json11.Task{Number: 6, Name: "feather"}
	seven := json11.Task{Number: 7, Name: "gold"}
	eight := json11.Task{Number: 8, Name: "hanger"}
	nine := json11.Task{Number: 9, Name: "inchworm"}
	ten := json11.Task{Number: 10, Name: "jellybean"}

	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		json11.PrintTasksToJSON(&buffer, one, two, three, four, five, six, seven, eight, nine, ten)
	}
}
