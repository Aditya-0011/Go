package routers

import (
	"net/http"
	"strings"

	"api/src/http/dtos"
	"api/src/http/handlers"
	"api/src/http/middlewares"
)

type MiddlewareOptions int8

const (
	None MiddlewareOptions = iota
	AuthRequired
	ValidationRequired
	AuthAndValidationRequired
)

func applyMiddlewares[T any](handler func(http.ResponseWriter, *http.Request, T), options MiddlewareOptions) http.HandlerFunc {
	switch options {
	case AuthRequired:
		return middlewares.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, *new(T))
		})

	case ValidationRequired:
		return middlewares.WithValidation(handler)

	case AuthAndValidationRequired:
		return middlewares.AuthMiddleware(middlewares.WithValidation(handler))

	default:
		return func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, *new(T))
		}
	}
}

func TodoRouter(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/todo/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/todo/")
		switch {
		case path == "getTodos" && r.Method == http.MethodGet:
			handler := func(w http.ResponseWriter, r *http.Request, _ struct{}) {
				h.GetTodos(w, r)
			}
			applyMiddlewares(handler, AuthRequired)(w, r)
		case path == "createTodo" && r.Method == http.MethodPost:
			handler := func(w http.ResponseWriter, r *http.Request, req dtos.CreateTodoRequest) {
				h.CreateTodo(w, r, req)
			}
			applyMiddlewares(handler, AuthAndValidationRequired)(w, r)
		case path == "updateTodo" && r.Method == http.MethodPost:
			handler := func(w http.ResponseWriter, r *http.Request, req dtos.UpdateTodoRequest) {
				h.UpdateTodo(w, r, req)
			}
			applyMiddlewares(handler, AuthAndValidationRequired)(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	return mux
}
