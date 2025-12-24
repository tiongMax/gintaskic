package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tiongMax/gintaskic/internal/handlers"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/tasks", handlers.GetTasksHandler)
		v1.GET("/tasks/:id", handlers.GetTaskByIDHandler)
		v1.POST("/tasks", handlers.CreateTaskHandler)
	}

	return r
}
