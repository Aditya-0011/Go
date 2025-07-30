package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"api/src/http/dtos"
	"api/src/models"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("select id, title, done from todos")
	if err != nil {
		http.Error(w, `{"msg":"failed to query todos"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.Id, &t.Title, &t.Done); err != nil {
			http.Error(w, `{"msg":"failed to scan todo"}`, http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos.GetTodosResponse{Todos: todos})
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request, req dtos.CreateTodoRequest) {
	_, err := h.DB.Exec("insert into todos (title, done) values (?, ?)", req.Title, req.Done)
	if err != nil {
		http.Error(w, `{"error":"failed to insert todo"}`, http.StatusInternalServerError)
		return
	}

	resp := dtos.TodoResponse{
		Msg: "todo created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request, req dtos.UpdateTodoRequest) {
	var currentTitle string
	var currentDone bool
	err := h.DB.QueryRow("SELECT title, done FROM todos WHERE id = ?", req.Id).Scan(&currentTitle, &currentDone)
	if err == sql.ErrNoRows {
		http.Error(w, `{"msg":"todo not found"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"msg":"failed to fetch todo"}`, http.StatusInternalServerError)
		return
	}

	newTitle := currentTitle
	newDone := currentDone

	if req.Title != nil {
		t := strings.TrimSpace(*req.Title)
		if t == "" {
			http.Error(w, `{"msg":"title cannot be empty"}`, http.StatusBadRequest)
			return
		}
		if len(t) > 200 {
			http.Error(w, `{"msg":"title too long (max 200 chars)"}`, http.StatusBadRequest)
			return
		}
		newTitle = t
	}

	if req.Done != nil {
		newDone = *req.Done
	}

	_, err = h.DB.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", newTitle, newDone, req.Id)
	if err != nil {
		http.Error(w, `{"msg":"failed to update todo"}`, http.StatusInternalServerError)
		return
	}

	resp := dtos.TodoResponse{
		Msg: "todo updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
