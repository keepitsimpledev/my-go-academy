package app

import (
	"errors"
	ds "go_academy/to-do-app/2/datastructure"
)

type TodoListAction struct {
	Action string `json:"Action"`
	Number int    `json:"Number"`
	Item   string `json:"Item"`
	Status string `json:"Status"`
}

func processAction(todoList *ds.TodoList, action TodoListAction) {
	switch action.Action {
	case "Add":
		todoList.Add(action.Item, action.Status)
	case "Update":
		err := todoList.Update(action.Number, action.Item, action.Status)
		panicIfErr(err)
	case "Delete":
		err := todoList.Delete(action.Number)
		panicIfErr(err)
	default:
		panic(errors.New("unexpected action: " + action.Action))
	}
}
