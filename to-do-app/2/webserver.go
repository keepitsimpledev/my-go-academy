package app

import (
	"encoding/json"
	"errors"
	"fmt"
	ds "go_academy/to-do-app/2/datastructure"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

const (
	jsonContentType    = "application/json"
	htmlTemplatePath   = "index.html"
	upgradedBufferSize = 1024
)

type RawTodoListAction struct {
	Action, Number, Item, Status string
}

type TodoListAction struct {
	Number               int
	Action, Item, Status string
}

type TodoListServer struct {
	http.Handler
	template *template.Template
	todoList ds.TodoList
}

func NewTodoListServer() (*TodoListServer, error) {
	todoListServer := new(TodoListServer)
	todoListServer.todoList = ds.NewList()
	htmlTemplate, err := template.ParseFiles(htmlTemplatePath)
	todoListServer.template = htmlTemplate

	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(todoListServer.init))
	router.Handle("/ws", http.HandlerFunc(todoListServer.webSocket))

	todoListServer.Handler = router

	return todoListServer, nil
}

//nolint:exhaustruct,gochecknoglobals
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  upgradedBufferSize,
	WriteBufferSize: upgradedBufferSize,
}

func (server *TodoListServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	defer conn.Close()

	updateList(conn, server)

	_, rawData, readMessageErr := conn.ReadMessage() // this waits for conn.send() from .html
	for readMessageErr == nil {
		// example: writes to conn.onmessage() in .html
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

func (server *TodoListServer) init(w http.ResponseWriter, _ *http.Request) {
	tmpl, parseErr := template.ParseFiles(htmlTemplatePath)

	if parseErr != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", parseErr.Error()), http.StatusInternalServerError)
		return
	}

	execErr := tmpl.Execute(w, nil)
	if execErr != nil {
		panic(execErr)
	}
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
	err := conn.WriteMessage(1, []byte(server.todoList.GetAll()))
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
