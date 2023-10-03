package routes

import (
	"context"

	"github.com/difmaj/gin-otel-api-demo/internal/pkg/server/routes/handlers"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Routes represents the routes of the application.
type Routes struct {
	Engine  *gin.Engine
	Handler *handlers.ToDoHandler
}

// New returns a new Routes instance.
// It also sets up the Gin engine and the CORS middleware.
func New() *Routes {

	cleanup := ConfigureOpenTelemetry()
	defer cleanup(context.Background())

	engine := gin.New()
	engine.Use(
		otelgin.Middleware("todo"),
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		}),
	)

	return &Routes{
		Engine:  engine,
		Handler: handlers.NewToDoHandler(),
	}
}

// Register registers all routes.
func (r *Routes) Register(group string) {
	api := r.Engine.Group(group)
	api.POST("todo", r.Handler.Create)
	api.GET("todo/:id", r.Handler.Get)
	api.GET("todo", r.Handler.GetAll)
	api.PUT("todo/:id", r.Handler.Update)
	api.DELETE("todo/:id", r.Handler.Delete)
}
