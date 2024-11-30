package app

import (
	"fmt"
	"net/http"
	"text/template"
)

func (server *TodoListServer) init(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only accepts GET", http.StatusBadRequest)
		return
	}

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
