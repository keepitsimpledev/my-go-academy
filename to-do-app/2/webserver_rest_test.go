package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	server := createNewServer()

	addAction := TodoListAction{
		Action: "Add",
		Number: 0,
		Item:   "wash dishes",
		Status: "Not Started",
	}
	request := actionToRequest(addAction, http.MethodPut)

	server.restHandler(createResponse(), request)

	got, getErr := server.todoList.Get(0)
	panicIfErr(getErr)

	want := "wash dishes - Not Started"
	assert(t, got, want)
}

func TestUpdate(t *testing.T) {
	server := createNewServer()

	server.todoList.Add("wash dishes", "Not Started")

	updateAction := TodoListAction{
		Action: "Update",
		Number: 1,
		Item:   "wash dishes",
		Status: "Complete",
	}
	updateRequest := actionToRequest(updateAction, http.MethodPut)

	server.restHandler(createResponse(), updateRequest)

	got, getErr := server.todoList.Get(0)
	panicIfErr(getErr)

	want := "wash dishes - Complete"
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

func assert(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got: '%s'. want: '%s'", got, want)
	}
}
