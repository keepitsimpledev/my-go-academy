package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

const (
	jsonContentType    = "application/json"
	upgradedBufferSize = 1024
)

//nolint:exhaustruct,gochecknoglobals
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  upgradedBufferSize,
	WriteBufferSize: upgradedBufferSize,
}

type RawTodoListAction struct {
	Action, Number, Item, Status string
}

type TodoListAction struct {
	Number               int
	Action, Item, Status string
}

func (server *TodoListServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	defer conn.Close()

	updateList(conn, server)

	_, rawData, readMessageErr := conn.ReadMessage() // this waits for conn.send() from .html
	for readMessageErr == nil {
		action, processInputErr := processInput(rawData)
		panicIfErr(processInputErr)

		switch action.Action {
		case "Add":
			server.todoList.Add(action.Item, action.Status)
		case "Update":
			err := server.todoList.Update(action.Number, action.Item, action.Status)
			panicIfErr(err)
		case "Delete":
			err := server.todoList.Delete(action.Number)
			panicIfErr(err)
		default:
			panic(errors.New("unexpected action: " + action.Action))
		}

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
