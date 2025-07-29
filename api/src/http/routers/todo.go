package routers

import (
	"net/http"
	"strings"

	"api/src/http/handlers"
)

func TodoRouter(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/todo/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/todo/")
		switch {
		case path == "getTodos" && r.Method == http.MethodGet:
			h.GetTodos(w, r)
		case path == "createTodo" && r.Method == http.MethodPost:
			h.CreateTodo(w, r)
		case path == "toggleTodo" && r.Method == http.MethodPost:
			h.ToggleTodo(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	return mux
}
