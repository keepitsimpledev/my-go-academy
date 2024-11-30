package app

import (
	"encoding/json"
	"go_academy/to-do-app/2/datastructure"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	server := createNewServer()

	task := ds.NewTask("wash dishes", "Not Started")
	addAction := TodoListAction{
		Action: "Add",
		Number: 0,
		Item:   task.GetItem(),
		Status: task.GetStatus(),
	}
	request := actionToRequest(addAction, http.MethodPut)

	server.restHandler(createResponse(), request)

	got, getErr := server.todoList.Get(0)
	panicIfErr(getErr)

	want := task
	assert(t, got, want)
}

func TestRead(t *testing.T) {
	server := createNewServer()
	task := ds.NewTask("wash dishes", "Not Started")

	server.todoList.Add(task.GetItem(), task.GetStatus())

	readRequest := httptest.NewRequest(http.MethodGet, "/task/1", nil)

	server.restHandler(createResponse(), readRequest)

	got, getErr := server.todoList.Get(0)
	panicIfErr(getErr)

	want := task
	assert(t, got, want)
}

func TestUpdate(t *testing.T) {
	server := createNewServer()
	initialTask := ds.NewTask("wash dishes", "Not Started")

	server.todoList.Add(initialTask.GetItem(), initialTask.GetStatus())

	updateTask := ds.NewTask(initialTask.GetItem(), "Complete")
	updateAction := TodoListAction{
		Action: "Update",
		Number: 1,
		Item:   updateTask.GetItem(),
		Status: updateTask.GetStatus(),
	}
	updateRequest := actionToRequest(updateAction, http.MethodPut)

	server.restHandler(createResponse(), updateRequest)

	got, getErr := server.todoList.Get(0)
	panicIfErr(getErr)

	want := updateTask
	assert(t, got, want)
}

func TestDelete(t *testing.T) {
	server := createNewServer()

	server.todoList.Add("wash dishes", "Not Started")

	deleteAction := TodoListAction{
		Action: "Delete",
		Number: 1,
		Item:   "",
		Status: "",
	}
	deleteRequest := actionToRequest(deleteAction, http.MethodDelete)

	server.restHandler(createResponse(), deleteRequest)

	got := server.todoList.GetAll()

	want := "To-Do list is empty\n"
	assert(t, got, want)
}

func createNewServer() *TodoListServer {
	server, err := NewTodoListServer()
	panicIfErr(err)

	return server
}

func actionToRequest(action TodoListAction, httpMethod string) *http.Request {
	body, marshalErr := json.Marshal(action)
	panicIfErr(marshalErr)

	bodyReader := strings.NewReader(string(body))

	return httptest.NewRequest(httpMethod, "/task", bodyReader)
}

func createResponse() *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	response.Header().Set("content-type", jsonContentType)

	return response
}

func assert(tb testing.TB, got, want any) {
	tb.Helper()

	if got != want {
		tb.Errorf("got: '%s'. want: '%s'", got, want)
	}
}
