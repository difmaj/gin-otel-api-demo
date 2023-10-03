package entities

import (
	"github.com/google/uuid"
)

// CreateToDoRequest defines the request body for creating a todo.
type CreateToDoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

// UpdateToDoRequest defines the request body for updating a todo.
type UpdateToDoRequest struct {
	ID        uuid.UUID `form:"id" binding:"required"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

// DeleteToDoRequest defines the request body for deleting a todo.
type DeleteToDoRequest struct {
	ID uuid.UUID `form:"id" binding:"required"`
}

// GetToDoRequest defines the request body for getting a todo.
type GetToDoRequest struct {
	ID uuid.UUID `form:"id" binding:"required"`
}

// ListToDoRequest defines the request body for listing todos.
type ListToDoRequest struct {
	Filters *ListToDoFilters `json:"filters"`
}

// ListToDoFilters defines the filters for listing todos.
type ListToDoFilters struct {
	Completed *bool `json:"completed"`
}
