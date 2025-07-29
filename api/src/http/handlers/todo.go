package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"api/src/models"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("select id, title, done from todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.DB.Exec("insert into todos (title, done) values (?, ?)", t.Title, t.Done)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	t.ID = strconv.FormatInt(id, 10)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (h *Handler) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")

	id, err := strconv.Atoi(query)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_, err = h.DB.Exec("update todos set done = not done WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
