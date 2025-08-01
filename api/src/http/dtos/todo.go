package dtos

import (
	"errors"
	"strings"

	"api/src/models"
)

type GetTodosResponse struct {
	Todos []models.Todo `json:"data"`
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (r CreateTodoRequest) Validate() error {
	r.Title = strings.TrimSpace(r.Title)
	if r.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(r.Title) < 3 || len(r.Title) > 100 {
		return errors.New("title must be between 3 and 100 characters")
	}
	return nil
}

type UpdateTodoRequest struct {
	Id    int     `json:"id"`
	Title *string `json:"title,omitempty"`
	Done  *bool   `json:"done,omitempty"`
}

func (r UpdateTodoRequest) Validate() error {
	if r.Id <= 0 {
		return errors.New("invalid id")
	}
	if r.Title != nil {
		t := strings.TrimSpace(*r.Title)
		if t == "" {
			return errors.New("title cannot be empty")
		}
		if len(t) > 100 {
			return errors.New("title too long (max 100 chars)")
		}
		*r.Title = t
	}
	return nil
}
