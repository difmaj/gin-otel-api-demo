package handlers

import (
	"net/http"

	"github.com/bluele/gcache"
	"github.com/difmaj/gin-otel-api-demo/internal/pkg/entities"
	"github.com/difmaj/gin-otel-api-demo/internal/pkg/repositories/todo"
	"github.com/difmaj/gin-otel-api-demo/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// ToDoHandler is the struct that contains the service instances.
type ToDoHandler struct {
	ToDoService *services.ToDoService
	Tracer      trace.Tracer
}

// NewToDoHandler returns a new Handler instance.
func NewToDoHandler() *ToDoHandler {
	return &ToDoHandler{
		Tracer: otel.Tracer("todo-handler"),
		ToDoService: services.NewToDoService(
			todo.NewCacheRepository(
				gcache.New(100).LRU().Build(),
			),
		),
	}
}

// Create creates a new todo.
func (h *ToDoHandler) Create(ctx *gin.Context) {
	_, span := h.Tracer.Start(
		ctx.Request.Context(),
		"Create",
	)
	defer span.End()

	var request entities.CreateToDoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.ToDoService.Create(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Get gets a todo by ID.
func (h *ToDoHandler) Get(ctx *gin.Context) {
	_, span := h.Tracer.Start(
		ctx.Request.Context(),
		"Get",
	)
	defer span.End()

	var request entities.GetToDoRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.ToDoService.Get(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetAll gets all todos.
func (h *ToDoHandler) GetAll(ctx *gin.Context) {
	_, span := h.Tracer.Start(
		ctx.Request.Context(),
		"GetAll",
	)
	defer span.End()

	var request entities.ListToDoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.ToDoService.GetAll(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Update updates a todo.
func (h *ToDoHandler) Update(ctx *gin.Context) {
	_, span := h.Tracer.Start(
		ctx.Request.Context(),
		"Update",
	)
	defer span.End()

	var request entities.UpdateToDoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.ToDoService.Update(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Delete deletes a todo by ID.
func (h *ToDoHandler) Delete(ctx *gin.Context) {
	_, span := h.Tracer.Start(
		ctx.Request.Context(),
		"Delete",
	)
	defer span.End()

	var request entities.DeleteToDoRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.ToDoService.Delete(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
