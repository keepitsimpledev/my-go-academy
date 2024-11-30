package todoprint_test

import (
	"bytes"
	todoprint "go_academy/to-do-app/1.16"
	"os"
	"testing"
)

func TestTodoType(t *testing.T) {
	wantTask := "wash dishes"
	wantStatus := "not started"
	got := todoprint.NewTodo(wantTask, wantStatus)

	if got.GetTask() != wantTask {
		t.Errorf("got: %s. want: %s", got.GetTask(), wantTask)
	}

	if got.GetStatus() != wantStatus {
		t.Errorf("got: %s. want: %s", got.GetStatus(), wantStatus)
	}
}

func TestPrintTasksAndStatuses(t *testing.T) {
	todos := []todoprint.Todo{
		todoprint.NewTodo("take a nap", "complete"),
		todoprint.NewTodo("do laundry", "in-progress"),
		todoprint.NewTodo("wash dishes", "not started"),
	}

	want := "take a nap - complete\ndo laundry - in-progress\nwash dishes - not started\n"

	buffer := bytes.Buffer{}
	todoprint.PrintTasksAndStatuses(&buffer, todos)
	got := buffer.String()

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}

func ExamplePrintTasksAndStatuses() {
	todos := []todoprint.Todo{
		todoprint.NewTodo("take a nap", "complete"),
		todoprint.NewTodo("do laundry", "in-progress"),
		todoprint.NewTodo("wash dishes", "not started"),
	}
	todoprint.PrintTasksAndStatuses(os.Stdout, todos)
	// Output:
	// take a nap - complete
	// do laundry - in-progress
	// wash dishes - not started
}

func BenchmarkPrintTasksAndStatuses(b *testing.B) {
	todos := []todoprint.Todo{
		todoprint.NewTodo("take a nap", "complete"),
		todoprint.NewTodo("do laundry", "in-progress"),
		todoprint.NewTodo("wash dishes", "not started"),
	}

	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		todoprint.PrintTasksAndStatuses(&buffer, todos)
	}
}
