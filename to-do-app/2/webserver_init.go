package app

import (
	"fmt"
	"net/http"
	"text/template"
)

func (server *TodoListServer) init(w http.ResponseWriter, _ *http.Request) {
	tmpl, parseErr := template.ParseFiles(htmlTemplatePath)

	if parseErr != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", parseErr.Error()), http.StatusInternalServerError)
		return
	}

	execErr := tmpl.Execute(w, nil)
	panicIfErr(execErr)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
