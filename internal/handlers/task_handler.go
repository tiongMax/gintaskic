package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiongMax/gintaskic/internal/model"
	"github.com/tiongMax/gintaskic/internal/repository/task"
)

// GetTasksHandler handles GET /tasks
func GetTasksHandler(c *gin.Context) {
	status := c.Query("status")
	tasks := task.GetAllTasks(status)
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByIDHandler handles GET /tasks/:id
func GetTaskByIDHandler(c *gin.Context) {
	id := c.Param("id")
	task, found := task.GetTaskByID(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTaskHandler(c *gin.Context) {
	var reqTask model.Task
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask := task.CreateTask(reqTask)
	c.JSON(http.StatusCreated, createdTask)
}
