package app

import (
	"encoding/json"
	ds "go_academy/to-do-app/2/datastructure"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

const testTimeoutMs = 1500

func TestWebsocket(t *testing.T) {
	// arrange
	todoListServer, newServerErr := NewTodoListServer()
	panicIfErr(newServerErr)

	testTodoListServer := httptest.NewServer(todoListServer)
	conn := mustDialWS(t, "ws"+strings.TrimPrefix(testTodoListServer.URL, "http")+"/ws")

	defer testTodoListServer.Close()
	defer conn.Close()

	// confirm initial state
	assertWebsocketGotMsg(t, conn, "To-Do list is empty\n")

	task := ds.NewTask("wash dishes", "Not Started")
	action := RawTodoListAction{
		Action: "Add",
		Number: "0",
		Item:   task.GetItem(),
		Status: task.GetStatus(),
	}
	rawAction, marshalErr := json.Marshal(action)
	panicIfErr(marshalErr)

	// act
	writeWSMessage(t, conn, string(rawAction))

	// assert
	assertMessage(t, &todoListServer.todoList, task)
	assertWebsocketGotMsg(t, conn, "To-Do list:\n1. wash dishes - Not Started\n")
}

func mustDialWS(tb testing.TB, url string) *websocket.Conn {
	tb.Helper()

	ws, response, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		tb.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	response.Body.Close()

	return ws
}

func writeWSMessage(tb testing.TB, conn *websocket.Conn, message string) {
	tb.Helper()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		tb.Fatalf("could not send message over ws connection %v", err)
	}
}

func assertMessage(tb testing.TB, todoList *ds.TodoList, want ds.Task) {
	tb.Helper()

	passed := retryUntil(testTimeoutMs*time.Millisecond, func() bool {
		got, err := todoList.Get(0)
		if err != nil {
			return false
		}

		return got == want
	})

	if !passed {
		tb.Errorf("wanted '%s' but never got", want)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}

	return false
}

func assertWebsocketGotMsg(tb testing.TB, conn *websocket.Conn, want string) {
	tb.Helper()

	_, msg, _ := conn.ReadMessage()
	if string(msg) != want {
		tb.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}
