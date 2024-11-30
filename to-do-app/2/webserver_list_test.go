package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	server, err := NewTodoListServer()
	panicIfErr(err)

	request := httptest.NewRequest(http.MethodGet, "/list", nil)
	response := httptest.NewRecorder()

	server.todoList.Add("wash dishes", "Not Started")

	server.list(response, request)

	rawGot, readErr := io.ReadAll(response.Body)
	panicIfErr(readErr)

	got := string(rawGot)

	want := "To-Do list:\n1. wash dishes - Not Started\n"
	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}
