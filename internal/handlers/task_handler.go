package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tiongMax/gintaskic/internal/model"
	"github.com/tiongMax/gintaskic/internal/repository/task"
)

// GetTasksHandler handles GET /tasks with search, pagination, and status filter
func GetTasksHandler(c *gin.Context) {
	status := c.Query("status")
	search := c.Query("search")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit < 1 {
		limit = 10
	}

	tasks, total, err := task.GetAllTasks(status, search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// GetTaskByIDHandler handles GET /tasks/:id
func GetTaskByIDHandler(c *gin.Context) {
	id := c.Param("id")
	foundTask, err := task.GetTaskByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if foundTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, foundTask)
}

func CreateTaskHandler(c *gin.Context) {
	var reqTask model.Task
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := task.CreateTask(reqTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTaskHandler handles PUT /tasks/:id (Full Update)
func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var reqTask model.Task
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert struct to map for update
	updateFields := map[string]interface{}{
		"title":  reqTask.Title,
		"status": reqTask.Status,
	}

	updatedTask, err := task.UpdateTask(id, updateFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// PatchTaskHandler handles PATCH /tasks/:id (Partial Update)
func PatchTaskHandler(c *gin.Context) {
	id := c.Param("id")

	// Bind to a map to get only supplied fields
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate allowed fields
	allowedFields := map[string]bool{"title": true, "status": true}
	cleanUpdates := make(map[string]interface{})

	for key, value := range updates {
		if allowedFields[key] {
			cleanUpdates[key] = value
		}
	}

	if len(cleanUpdates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return
	}

	updatedTask, err := task.UpdateTask(id, cleanUpdates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTaskHandler handles DELETE /tasks/:id
func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	err := task.DeleteTask(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// BulkDeleteTasksHandler handles DELETE /tasks (e.g., clear completed)
func BulkDeleteTasksHandler(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status query parameter is required for bulk delete"})
		return
	}

	count, err := task.DeleteTasksByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Bulk delete successful",
		"deleted_count": count,
	})
}
