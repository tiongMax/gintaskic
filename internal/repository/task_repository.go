package repository

import (
	"github.com/tiongMax/gintaskic/internal/model"
)

// GetAllTasks retrieves tasks, optionally filtered by status
func GetAllTasks(status string) []model.Task {
	if status == "" {
		return model.MockTasks
	}

	// Filter tasks (avoid modifying the original slice in-place if we were pointing to it directly)
	// Here we return a new slice
	var filtered []model.Task
	for _, task := range model.MockTasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// GetTaskByID finds a task by its ID
func GetTaskByID(id string) (*model.Task, bool) {
	for _, task := range model.MockTasks {
		if task.ID == id {
			return &task, true
		}
	}
	return nil, false
}
