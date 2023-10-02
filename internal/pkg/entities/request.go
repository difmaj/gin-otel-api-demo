package entities

import (
	"github.com/google/uuid"
)

// CreateToDoRequest is the struct that defines the request body for creating a new todo.
type CreateToDoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

// UpdateToDoRequest is the struct that defines the request body for updating a todo.
type UpdateToDoRequest struct {
	ID        uuid.UUID `form:"id" binding:"required"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

// DeleteToDoRequest is the struct that defines the request body for deleting a todo.
type DeleteToDoRequest struct {
	ID uuid.UUID `form:"id" binding:"required"`
}

// GetToDoRequest is the struct that defines the request body for getting a todo.
type GetToDoRequest struct {
	ID uuid.UUID `form:"id" binding:"required"`
}

// ListToDoRequest is the struct that defines the request body for getting all todos.
type ListToDoRequest struct {
	Filters *ListToDoFilters `json:"filters"`
}

// ListToDoFilters is the struct that defines the filters for getting all todos.
type ListToDoFilters struct {
	Completed *bool `json:"completed"`
}
