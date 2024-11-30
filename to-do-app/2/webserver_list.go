package app

import "net/http"

func (server *TodoListServer) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only accepts GET", http.StatusBadRequest)
		return
	}

	all := server.todoList.GetAll()
	_, err := w.Write([]byte(all))
	panicIfErr(err)

	w.WriteHeader(http.StatusAccepted)
}
