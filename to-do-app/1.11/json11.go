package json11

import (
	"encoding/json"
	"fmt"
	"io"
)

type Task struct {
	Number int
	Name   string
}

func PrintTasksToJSON(writer io.Writer, things ...Task) {
	numberedThingsJSON, err := json.Marshal(things)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(writer, string(numberedThingsJSON))
}
