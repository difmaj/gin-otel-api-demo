package entities

import "errors"

// application errors
var (
	ErrFailureToRecoverData = errors.New("could not recover data")
	ErrToDoNotFound         = errors.New("todo not found")
	ErrFailedToRemoveToDo   = errors.New("could not remove todo")
	ErrMissingTitle         = errors.New("missing title")
)
