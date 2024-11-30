package ds

import (
	"errors"
	"fmt"
	"sync"
)

const oobErrorMessage = "index out-of-bounds"

type TodoList struct {
	mu    sync.Mutex
	tasks []task
}

func NewList() TodoList {
	var tasks []task

	return TodoList{
		mu:    sync.Mutex{},
		tasks: tasks,
	}
}

func (tl *TodoList) Add(item, status string) {
	tl.mu.Lock()

	task := task{item, status}
	tl.tasks = append(tl.tasks, task)

	tl.mu.Unlock()
}

func (tl *TodoList) Get(index int) (string, error) {
	tl.mu.Lock()

	var task string

	var err error

	if index >= len(tl.tasks) {
		err = errors.New(oobErrorMessage)
	} else {
		task = tl.tasks[index].String()
	}

	tl.mu.Unlock()

	return task, err
}

func (tl *TodoList) GetAll() string {
	tl.mu.Lock()

	var tasks string

	if len(tl.tasks) == 0 {
		tasks = "To-Do list is empty\n"
	} else {
		tasks = "To-Do list:\n"
		for i := 0; i < len(tl.tasks); i++ {
			tasks += fmt.Sprintf("%d. %s\n", i+1, tl.tasks[i].String())
		}
	}

	tl.mu.Unlock()

	return tasks
}

func (tl *TodoList) Update(index int, item, status string) error {
	tl.mu.Lock()

	var err error

	if index >= len(tl.tasks) || index < 0 {
		err = errors.New(oobErrorMessage)
	} else {
		tl.tasks[index].item = item
		tl.tasks[index].status = status
	}

	tl.mu.Unlock()

	return err
}

func (tl *TodoList) Delete(index int) error {
	tl.mu.Lock()

	var err error

	if index >= len(tl.tasks) {
		err = errors.New(oobErrorMessage)
	} else {
		tl.tasks = append(tl.tasks[0:index], tl.tasks[index+1:]...)
	}

	tl.mu.Unlock()

	return err
}
