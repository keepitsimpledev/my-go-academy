package readandprint_test

import (
	"fmt"
	readandprint "go_academy/to-do-app/1.13"
	"slices"
	"testing"
	"testing/fstest"
)

func TestReadJSONTaskFiles(t *testing.T) {
	// arrange
	content := `[{"Priority": 1, "Task": "laundry"}, {"Priority": 2, "Task": "dishes"}, {"Priority": 3, "Task": "taxes"}]`
	fs := fstest.MapFS{
		"tasks.md": {Data: []byte(content)},
	}

	// act
	got, _ := readandprint.ReadJSONTaskFiles(fs)

	// assert
	want := []readandprint.Task{
		{Priority: 1, Task: "laundry"},
		{Priority: 2, Task: "dishes"},
		{Priority: 3, Task: "taxes"},
	}
	if !slices.Equal(got, want) {
		t.Errorf("got:\n%v\n\nwant:\n%v", got, want)
	}
}

func ExampleReadJSONTaskFiles() {
	// set up sample .json file with tasks
	sampleContent := `[{"Priority": 1, "Task": "finish homework"}, {"Priority": 2, "Task": "defrost chicken"}]`
	fs := fstest.MapFS{
		"sample-tasks-file.json": {Data: []byte(sampleContent)},
	}
	tasks, _ := readandprint.ReadJSONTaskFiles(fs)
	fmt.Println(tasks)
	// Output: [{1 finish homework} {2 defrost chicken}]
}

func BenchmarkReadJSONTaskFiles(b *testing.B) {
	content := `[{"Priority": 1, "Task": "study"}, {"Priority": 2, "Task": "clean"}, {"Priority": 3, "Task": "shower"},
		{"Priority": 4, "Task": "nap"}, {"Priority": 5, "Task": "manage"}, {"Priority": 6, "Task": "call"},
		{"Priority": 7, "Task": "frame"}, {"Priority": 8, "Task": "repair"}, {"Priority": 9, "Task": "maintain"}
		]`
	fs := fstest.MapFS{
		"sample-tasks-file.json": {Data: []byte(content)},
	}

	for i := 0; i < b.N; i++ {
		_, err := readandprint.ReadJSONTaskFiles(fs)
		if err != nil {
			panic(err)
		}
	}
}
