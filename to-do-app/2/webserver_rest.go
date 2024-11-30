package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (server *TodoListServer) restHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		getExec(responseWriter, request, server)
	} else if request.Method == http.MethodPut || request.Method == http.MethodDelete {
		restExec(responseWriter, request, server)
	} else {
		http.Error(responseWriter, "only accepts PUT and DELETE", http.StatusBadRequest)
	}
}

func getExec(responseWriter http.ResponseWriter, request *http.Request, server *TodoListServer) {
	path := strings.Split(request.RequestURI, "/")
	if len(path) < 2 || path[len(path)-2] != "task" {
		http.Error(responseWriter, fmt.Sprintf("invalid path: %s", request.RequestURI), http.StatusBadRequest)
		return
	}

	index, atoiErr := strconv.Atoi(path[len(path)-1])
	if atoiErr != nil {
		http.Error(responseWriter, fmt.Sprintf("parse error: %s", atoiErr.Error()), http.StatusInternalServerError)
		return
	}

	task, getErr := server.todoList.Get(index - 1)
	if getErr != nil {
		http.Error(responseWriter, fmt.Sprintf("failed to retrieve task: %s", getErr.Error()), http.StatusInternalServerError)
		return
	}

	marshallableTask := struct{ Item, Status string }{
		Item:   task.GetItem(),
		Status: task.GetStatus(),
	}

	taskBytes, marshalErr := json.Marshal(&marshallableTask)
	if marshalErr != nil {
		http.Error(responseWriter, fmt.Sprintf("failed to parse task: %s", marshalErr.Error()),
			http.StatusInternalServerError)
		return
	}

	_, writeErr := responseWriter.Write(taskBytes)
	if writeErr != nil {
		http.Error(responseWriter, fmt.Sprintf("write failure: %s", writeErr.Error()), http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusAccepted)
}

func restExec(responseWriter http.ResponseWriter, request *http.Request, server *TodoListServer) {
	rawBody, readErr := io.ReadAll(request.Body)
	panicIfErr(readErr)

	var action TodoListAction
	unmarshalErr := json.Unmarshal(rawBody, &action)
	panicIfErr(unmarshalErr)

	responseCode := -1

	var err error

	switch action.Action {
	case "Add":
		server.todoList.Add(action.Item, action.Status)
	case "Update":
		err = server.todoList.Update(action.Number-1, action.Item, action.Status)
	case "Delete":
		err = server.todoList.Delete(action.Number - 1)
	default:
		err = fmt.Errorf("unexpected action: %s", action.Action)
		responseCode = http.StatusBadRequest
	}

	if err != nil {
		if responseCode == -1 {
			responseCode = http.StatusInternalServerError
		}

		http.Error(responseWriter, fmt.Sprintf("error: %s", err.Error()), responseCode)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}
