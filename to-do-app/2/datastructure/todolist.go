package ds

import (
	"errors"
	"fmt"
)

const oobErrorMessage = "index out-of-bounds"

type TodoList struct {
	tasks []task
}

func NewList() TodoList {
	var tasks []task
	return TodoList{tasks: tasks}
}

func (tl *TodoList) Add(item, status string) {
	task := task{item, status}
	tl.tasks = append(tl.tasks, task)
}

func (tl *TodoList) Get(index int) (string, error) {
	if index >= len(tl.tasks) {
		return "", errors.New(oobErrorMessage)
	}

	return tl.tasks[index].String(), nil
}

func (tl *TodoList) GetAll() string {
	if len(tl.tasks) == 0 {
		return "To-Do list is empty\n"
	}

	all := "To-Do list:\n"
	for i := 0; i < len(tl.tasks); i++ {
		all += fmt.Sprintf("%d. %s\n", i+1, tl.tasks[i].String())
	}

	return all
}

func (tl *TodoList) Update(index int, item, status string) error {
	if index >= len(tl.tasks) || index < 0 {
		return errors.New(oobErrorMessage)
	}

	tl.tasks[index].item = item
	tl.tasks[index].status = status

	return nil
}

func (tl *TodoList) Delete(index int) error {
	if index >= len(tl.tasks) {
		return errors.New(oobErrorMessage)
	}

	tl.tasks = append(tl.tasks[0:index], tl.tasks[index+1:]...)

	return nil
}
