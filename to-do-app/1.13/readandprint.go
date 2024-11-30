package readandprint

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
)

type Task struct {
	Priority int    `json:"Priority"`
	Task     string `json:"Task"`
}

func ReadJSONTaskFiles(fileSystem fs.FS) ([]Task, error) {
	directoryEntries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, fmt.Errorf("io/fs.ReadDir error: %v", err)
	}

	var tasks []Task

	for _, directoryEntry := range directoryEntries {
		newTasks := parseJSONFile(fileSystem, directoryEntry.Name())
		tasks = append(tasks, newTasks...)
	}

	return tasks, nil
}

func parseJSONFile(fileSystem fs.FS, filename string) []Task {
	file, openErr := fileSystem.Open(filename)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}

	var tasks []Task

	unmarshalErr := json.Unmarshal(buffer.Bytes(), &tasks)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return tasks
}
