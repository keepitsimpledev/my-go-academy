package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test is essentially a test for this request:
// /batch?actions=[{"Action"%3A"Add"%2C"Number"%3A0%2C"Item"%3A"wash+dishes"%2C"Status"%3A"Not+Started"}
// %2C{"Action"%3A"Add"%2C"Number"%3A0%2C"Item"%3A"clean+room"%2C"Status"%3A"Not+Started"}]
func TestWebserverBatch(t *testing.T) {
	todoListServer, newListErr := NewTodoListServer()
	panicIfErr(newListErr)

	actions := []TodoListAction{
		{
			Action: "Add",
			Number: 0,
			Item:   "wash dishes",
			Status: "Not Started",
		},
		{
			Action: "Add",
			Number: 0,
			Item:   "clean room",
			Status: "Not Started",
		},
	}
	stringifiedactions, marshalErr := json.Marshal(actions)
	panicIfErr(marshalErr)

	request := httptest.NewRequest(http.MethodGet, "/batch", nil)

	queryCopy := request.URL.Query()
	queryCopy.Add("actions", string(stringifiedactions))
	request.URL.RawQuery = queryCopy.Encode()

	response := httptest.NewRecorder()

	todoListServer.batch(response, request)

	got := todoListServer.todoList.GetAll()

	assertContains(t, got, "To-Do list:\n")
	assertContains(t, got, "1. ")
	assertContains(t, got, "wash dishes - Not Started\n")
	assertContains(t, got, "2. ")
	assertContains(t, got, "clean room - Not Started\n")
}

func assertContains(tb testing.TB, got, want string) {
	tb.Helper()

	if !strings.Contains(got, want) {
		tb.Errorf("got:\n%s\n\nwant a string containing:\n%s", got, want)
	}
}
