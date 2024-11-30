package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

const upgradedBufferSize = 1024

//nolint:exhaustruct,gochecknoglobals
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  upgradedBufferSize,
	WriteBufferSize: upgradedBufferSize,
}

type RawTodoListAction struct {
	Action string `json:"Action"`
	Number string `json:"Number"`
	Item   string `json:"Item"`
	Status string `json:"Status"`
}

func (server *TodoListServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	defer conn.Close()

	updateList(conn, server)

	_, rawData, readMessageErr := conn.ReadMessage() // this waits for conn.send() from .html
	for readMessageErr == nil {
		action, processInputErr := processInput(rawData)
		panicIfErr(processInputErr)

		processAction(&server.todoList, action)

		updateList(conn, server)

		_, rawData, readMessageErr = conn.ReadMessage()
	}
	panicIfErr(readMessageErr)
}

func processInput(rawData []byte) (TodoListAction, error) {
	var rawAction RawTodoListAction

	dataString := string(rawData)
	escapedDataString := strings.ReplaceAll(dataString, "\"", "\\\"")
	quotedDataString := `"` + escapedDataString + `"`

	unquotedString, unquoteErr := strconv.Unquote(quotedDataString)
	if unquoteErr != nil {
		return TodoListAction{}, fmt.Errorf("processInput error: %v", unquoteErr)
	}

	unmarshalErr := json.Unmarshal([]byte(unquotedString), &rawAction)
	if unmarshalErr != nil {
		return TodoListAction{}, fmt.Errorf("processInput error: %v", unmarshalErr)
	}

	number, atoiErr := strconv.Atoi(rawAction.Number)
	if atoiErr != nil {
		return TodoListAction{}, fmt.Errorf("processInput error: %v", atoiErr)
	}

	return TodoListAction{
		Action: rawAction.Action,
		Number: number - 1,
		Item:   rawAction.Item,
		Status: rawAction.Status,
	}, nil
}

func updateList(conn *websocket.Conn, server *TodoListServer) {
	err := conn.WriteMessage(1, []byte(server.todoList.GetAll())) // writes to conn.onmessage() in .html
	panicIfErr(err)
}
