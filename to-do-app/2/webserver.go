package app

import (
	"fmt"
	ds "go_academy/to-do-app/2/datastructure"
	"net/http"
	"text/template"
)

const htmlTemplatePath = "index.html"

type TodoListServer struct {
	http.Handler
	template *template.Template
	todoList ds.TodoList
}

func NewTodoListServer() (*TodoListServer, error) {
	todoListServer := new(TodoListServer)
	todoListServer.todoList = ds.NewList()
	htmlTemplate, err := template.ParseFiles(htmlTemplatePath)
	todoListServer.template = htmlTemplate

	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(todoListServer.init))
	router.Handle("/ws", http.HandlerFunc(todoListServer.webSocket))
	router.Handle("/batch", http.HandlerFunc(todoListServer.batch))

	todoListServer.Handler = router

	return todoListServer, nil
}
