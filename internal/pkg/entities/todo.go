package entities

import "github.com/google/uuid"

// ToDo defines the todo entity.
type ToDo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}
