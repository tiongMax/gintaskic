package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiongMax/gintaskic/internal/handlers"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", handlers.GetTasksHandler)
		v1.GET("/tasks/:id", handlers.GetTaskByIDHandler)
		v1.POST("/tasks", handlers.CreateTaskHandler)
	}

	router.Run(":8082")
}
