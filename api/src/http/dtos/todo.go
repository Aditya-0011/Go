package dtos

import (
	"errors"
	"strings"

	"api/src/models"
)

type TodoResponse struct {
	Msg string `json:"msg"`
}

type GetTodosResponse struct {
	Todos []models.Todo `json:"todos"`
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (c CreateTodoRequest) Validate() error {
	c.Title = strings.TrimSpace(c.Title)
	if c.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(c.Title) < 3 || len(c.Title) > 100 {
		return errors.New("title must be between 3 and 100 characters")
	}
	return nil
}

type UpdateTodoRequest struct {
	Id    int     `json:"id"`
	Title *string `json:"title,omitempty"`
	Done  *bool   `json:"done,omitempty"`
}

func (u UpdateTodoRequest) Validate() error {
	if u.Id <= 0 {
		return errors.New("invalid id")
	}
	if u.Title != nil {
		t := strings.TrimSpace(*u.Title)
		if t == "" {
			return errors.New("title cannot be empty")
		}
		if len(t) > 100 {
			return errors.New("title too long (max 100 chars)")
		}
	}
	return nil
}
