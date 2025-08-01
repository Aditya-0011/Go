package routers

import (
	"database/sql"
	"net/http"
	"strings"

	"api/src/http/dtos"
	"api/src/http/handlers"
	"api/src/http/middlewares"
)

func TodoRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	h := &handlers.TodoHandler{DB: db}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/todo/")
		switch {
		case path == "getTodos" && r.Method == http.MethodGet:
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.GetTodos(w, r)
			})
			middlewares.AddMiddlewares(handler, middlewares.AuthMiddleware).ServeHTTP(w, r)
		case path == "createTodo" && r.Method == http.MethodPost:
			handler := func(w http.ResponseWriter, r *http.Request, req dtos.CreateTodoRequest) {
				h.CreateTodo(w, r, req)
			}
			middlewares.AddMiddlewares(
				middlewares.ValidationMiddleware(handler),
				middlewares.AuthMiddleware,
			).ServeHTTP(w, r)
		case path == "updateTodo" && r.Method == http.MethodPost:
			handler := func(w http.ResponseWriter, r *http.Request, req dtos.UpdateTodoRequest) {
				h.UpdateTodo(w, r, req)
			}
			middlewares.AddMiddlewares(
				middlewares.ValidationMiddleware(handler),
				middlewares.AuthMiddleware,
			).ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	return mux
}
