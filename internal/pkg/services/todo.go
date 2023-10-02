package services

import (
	"context"

	"github.com/difmaj/gin-otel-api-demo/internal/pkg/entities"
)

// ToDoRepository is the interface that defines the methods that the Repository struct must implement.
type ToDoRepository interface {
	// Create creates a new todo.
	Create(context.Context, *entities.CreateToDoRequest) (*entities.CreateToDoResponse, error)
	// Get gets a todo by its ID.
	Get(context.Context, *entities.GetToDoRequest) (*entities.GetToDoResponse, error)
	// List gets all todos.
	List(context.Context, *entities.ListToDoRequest) (*entities.GetToDosResponse, error)
	// Update updates a todo.
	Update(context.Context, *entities.UpdateToDoRequest) (*entities.UpdateToDoResponse, error)
	// Delete deletes a todo by its ID.
	Delete(context.Context, *entities.DeleteToDoRequest) (*entities.DeleteToDoResponse, error)
}

// ToDoService is the struct that implements the UseCase interface.
type ToDoService struct {
	repository ToDoRepository
}

// NewToDoService returns a new Service instance.
func NewToDoService(repository ToDoRepository) *ToDoService {
	return &ToDoService{repository: repository}
}

// Create creates a new todo.
func (s *ToDoService) Create(ctx context.Context, request *entities.CreateToDoRequest) (*entities.CreateToDoResponse, error) {
	return s.repository.Create(ctx, request)
}

// Get gets a todo by its ID.
func (s *ToDoService) Get(ctx context.Context, request *entities.GetToDoRequest) (*entities.GetToDoResponse, error) {
	return s.repository.Get(ctx, request)
}

// GetAll gets all todos.
func (s *ToDoService) GetAll(ctx context.Context, request *entities.ListToDoRequest) (*entities.GetToDosResponse, error) {
	return s.repository.List(ctx, request)
}

// Update updates a todo.
func (s *ToDoService) Update(ctx context.Context, request *entities.UpdateToDoRequest) (*entities.UpdateToDoResponse, error) {
	return s.repository.Update(ctx, request)
}

// Delete deletes a todo by its ID.
func (s *ToDoService) Delete(ctx context.Context, request *entities.DeleteToDoRequest) (*entities.DeleteToDoResponse, error) {
	return s.repository.Delete(ctx, request)
}
