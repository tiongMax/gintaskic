package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tiongMax/gintaskic/internal/config"
	"github.com/tiongMax/gintaskic/internal/handlers"
	"github.com/tiongMax/gintaskic/internal/middleware"
)

func SetupRouter(r *gin.Engine, cfg *config.Config) *gin.Engine {
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/tasks", handlers.GetTasksHandler)
		v1.GET("/tasks/:id", handlers.GetTaskByIDHandler)
		v1.POST("/tasks", handlers.CreateTaskHandler)
		v1.PUT("/tasks/:id", handlers.UpdateTaskHandler)
		v1.PATCH("/tasks/:id", handlers.PatchTaskHandler) // New
		v1.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
		v1.DELETE("/tasks", handlers.BulkDeleteTasksHandler) // New
	}

	return r
}
