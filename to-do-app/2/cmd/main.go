package main

import (
	"go_academy/to-do-app/2"
	"log"
	"net/http"
)

func main() {
	server, err := app.NewTodoListServer()
	if err != nil {
		panic(err)
	}

	//nolint:gosec
	log.Fatal(http.ListenAndServe(":5000", server))
}
