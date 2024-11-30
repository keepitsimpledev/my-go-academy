package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (server *TodoListServer) restHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut || request.Method == http.MethodDelete {
		restExec(responseWriter, request, server)
	} else {
		http.Error(responseWriter, "only accepts PUT and DELETE", http.StatusBadRequest)
	}
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
