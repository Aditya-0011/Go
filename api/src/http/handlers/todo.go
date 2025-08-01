package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"api/src/http/dtos"
	"api/src/models"
	"api/src/utils"
)

type TodoHandler struct {
	DB *sql.DB
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("select id, title, done from todos")
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Msg: "Failed to query todos"})
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.Id, &t.Title, &t.Done); err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Msg: "Failed to scan todo"})
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos.GetTodosResponse{Todos: todos})
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request, req dtos.CreateTodoRequest) {

	_, err := h.DB.Exec("insert into todos (title, done) values (?, ?)", req.Title, req.Done)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Msg: "Failed to insert todo"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct{}{})
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, req dtos.UpdateTodoRequest) {
	var currentTitle string
	var currentDone bool
	err := h.DB.QueryRow("SELECT title, done FROM todos WHERE id = ?", req.Id).Scan(&currentTitle, &currentDone)
	if err == sql.ErrNoRows {
		utils.WriteJSON(w, http.StatusNotFound, utils.ErrorResponse{Msg: "Todo not found"})
		return
	} else if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Msg: "Failed to fetch todo"})
		return
	}

	newTitle := currentTitle
	newDone := currentDone

	if req.Done != nil {
		newDone = *req.Done
	}

	_, err = h.DB.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", newTitle, newDone, req.Id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Msg: "Failed to update todo"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct{}{})
}
