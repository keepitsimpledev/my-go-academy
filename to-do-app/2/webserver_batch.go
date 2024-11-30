package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (server *TodoListServer) batch(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	params, parseErr := url.ParseQuery(urlString)
	panicIfErr(parseErr)

	for rawParam := range params {
		param := cleanupParamName(rawParam)
		if param == "actions" {
			rawActions := params.Get(rawParam)

			var actions []TodoListAction
			unmarshalError := json.Unmarshal([]byte(rawActions), &actions)
			panicIfErr(unmarshalError)

			actionChannel := make(chan TodoListAction)

			for _, action := range actions {
				go func(action TodoListAction) {
					actionChannel <- action
				}(action)
			}

			for i := 0; i < len(actions); i++ {
				action := <-actionChannel
				processAction(&server.todoList, action)
			}
		}
	}

	fmt.Fprint(w, "processed.")
}

func cleanupParamName(paramName string) string {
	paramName, _ = strings.CutPrefix(paramName, "/")
	paramName, _ = strings.CutPrefix(paramName, "batch")
	paramName, _ = strings.CutPrefix(paramName, "?")

	return paramName
}
