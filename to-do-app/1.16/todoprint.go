package todoprint

import (
	"fmt"
	"io"
	"sync"
)

const numTodoProperties = 2

func PrintTasksAndStatuses(writer io.Writer, todos []Todo) {
	var waitGroup sync.WaitGroup

	muTask := sync.Mutex{}
	muStatus := sync.Mutex{}

	waitGroup.Add(numTodoProperties)
	muStatus.Lock()

	go func(todos []Todo) {
		for i := 0; i < len(todos); i++ {
			muTask.Lock()
			fmt.Fprintf(writer, "%s - ", todos[i].GetTask())
			muStatus.Unlock()
		}
		waitGroup.Done()
	}(todos)

	go func(todos []Todo) {
		for i := 0; i < len(todos); i++ {
			muStatus.Lock()
			fmt.Fprintf(writer, "%s\n", todos[i].GetStatus())
			muTask.Unlock()
		}
		waitGroup.Done()
	}(todos)

	waitGroup.Wait()
}
