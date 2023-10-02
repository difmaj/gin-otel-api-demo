package todo

import (
	"context"
	"encoding/json"

	"github.com/bluele/gcache"
	"github.com/difmaj/gin-otel-api-demo/internal/pkg/entities"
	"github.com/difmaj/gin-otel-api-demo/internal/pkg/services"
	"github.com/google/uuid"
)

// CacheRepository is the struct that implements the todo.CacheRepository interface.
type CacheRepository struct {
	cache gcache.Cache
}

// NewCacheRepository returns a new Repository instance.
func NewCacheRepository(cache gcache.Cache) services.ToDoRepository {
	return &CacheRepository{
		cache: cache,
	}
}

// recoverData is a helper function to recover data from the cache.
func (r *CacheRepository) recoverData(data interface{}) (*entities.ToDo, error) {
	if data == nil {
		return nil, entities.ErrToDoNotFound
	}

	value, valid := data.([]byte)
	if !valid {
		return nil, entities.ErrFailureToRecoverData
	}

	recovered := new(entities.ToDo)
	err := json.Unmarshal(value, recovered)
	if err != nil {
		return nil, err
	}
	return recovered, nil
}

// Create creates a new ToDo item in the cache.
func (r *CacheRepository) Create(ctx context.Context, request *entities.CreateToDoRequest) (*entities.CreateToDoResponse, error) {

	if request.Title == "" {
		return nil, entities.ErrMissingTitle
	}

	newID := uuid.New()

	todo := new(entities.ToDo)
	todo.Completed = request.Completed
	todo.Title = request.Title
	todo.ID = newID

	b, err := json.Marshal(todo)
	if err != nil {
		return nil, err
	}

	err = r.cache.Set(newID, b)
	if err != nil {
		return nil, err
	}

	response, err := r.Get(ctx, &entities.GetToDoRequest{ID: newID})
	if err != nil {
		return nil, err
	}

	return &entities.CreateToDoResponse{
		Data: response.Data,
	}, nil
}

// Get retrieves a ToDo item from the cache by ID.
func (r *CacheRepository) Get(ctx context.Context, request *entities.GetToDoRequest) (*entities.GetToDoResponse, error) {
	response, err := r.cache.Get(request.ID)
	if err != nil {
		return nil, err
	}
	recovered, err := r.recoverData(response)
	if err != nil {
		return nil, err
	}
	return &entities.GetToDoResponse{
		Data: recovered,
	}, nil
}

// List retrieves all ToDos from the cache.
func (r *CacheRepository) List(ctx context.Context, request *entities.ListToDoRequest) (*entities.GetToDosResponse, error) {

	todos := make([]*entities.ToDo, 0, r.cache.Len(true))
	caches := r.cache.GetALL(true)

	for _, value := range caches {
		recovered, err := r.recoverData(value)
		if err != nil {
			return nil, err
		}

		if request != nil && request.Filters != nil {
			if request.Filters.Completed != nil &&
				*request.Filters.Completed != recovered.Completed {
				continue
			}
		}
		todos = append(todos, recovered)
	}
	return &entities.GetToDosResponse{
		Data: todos,
	}, nil
}

// Update updates a todo.
func (r *CacheRepository) Update(ctx context.Context, request *entities.UpdateToDoRequest) (*entities.UpdateToDoResponse, error) {

	todo, err := r.Get(ctx, &entities.GetToDoRequest{ID: request.ID})
	if err != nil {
		return nil, err
	}

	todo.Data.Title = request.Title
	todo.Data.Completed = request.Completed

	b, err := json.Marshal(todo)
	if err != nil {
		return nil, err
	}

	err = r.cache.Set(request.ID, b)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateToDoResponse{
		Data: todo.Data,
	}, nil
}

// Delete deletes a todo by its ID.
func (r *CacheRepository) Delete(ctx context.Context, request *entities.DeleteToDoRequest) (*entities.DeleteToDoResponse, error) {
	removed := r.cache.Remove(request.ID)
	if !removed {
		return nil, entities.ErrFailedToRemoveToDo
	}

	return &entities.DeleteToDoResponse{
		Success: true,
	}, nil
}
