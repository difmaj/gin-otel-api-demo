package entities

import "errors"

// application errors
var (
	// ErrFailureToRecoverToDoData is returned when the data could not be recovered.
	ErrFailureToRecoverToDoData = errors.New("could not recover data")
	// ErrToDoNotFound is returned when the todo is not found.
	ErrToDoNotFound = errors.New("todo not found")
	// ErrFailedToRemoveToDo is returned when the todo could not be removed.
	ErrFailedToRemoveToDo = errors.New("could not remove todo")
	// ErrMissingToDoTitle is returned when the title is missing.
	ErrMissingToDoTitle = errors.New("missing todo title")
)
