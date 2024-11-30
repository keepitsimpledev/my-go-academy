package main

import (
	"go_academy/to-do-app/2"
	"log"
	"net/http"
	"os"
	"sync"
)

const numberOfThreads = 2

func main() {
	var wg sync.WaitGroup

	wg.Add(numberOfThreads)

	go initWebserver()
	initCmdLine(&wg)

	wg.Wait() // unreachable, but keeps webserver alive
}

func initWebserver() {
	server, err := app.NewTodoListServer()
	if err != nil {
		panic(err)
	}

	//nolint:gosec
	log.Fatal(http.ListenAndServe(":5000", server))
}

func initCmdLine(wg *sync.WaitGroup) {
	app.Run(os.Stdin, os.Stdout)
	wg.Done()
}
