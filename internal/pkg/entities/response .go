package entities

// Error defines the error response.
type Error struct {
	Message  string `json:"message"`
	Metadata any    `json:"metadata"`
}

// Response defines the response wrapper.
type Response[T any] struct {
	Error   *Error `json:"error,omitempty"`
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
}

// CreateToDoResponse defines the response for creating a new todo.
type CreateToDoResponse Response[*ToDo]

// UpdateToDoResponse defines the response for updating a todo.
type UpdateToDoResponse Response[*ToDo]

// DeleteToDoResponse defines the response for deleting a todo.
type DeleteToDoResponse struct {
	Error   *Error `json:"error,omitempty"`
	Success bool   `json:"success"`
}

// GetToDoResponse defines the response for getting a todo.
type GetToDoResponse Response[*ToDo]

// GetToDosResponse defines the response for getting all todos.
type GetToDosResponse Response[[]*ToDo]
