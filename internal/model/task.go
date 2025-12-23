package model

import (
	"time"
)

const (
	TaskStatusPending    = "pending"
	TaskStatusCompleted  = "completed"
	TaskStatusInProgress = "in_progress"
	TaskStatusCancelled  = "cancelled"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" binding:"required,min=3"`
	Status    string    `json:"status" binding:"required,oneof=pending completed in_progress cancelled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var MockTask = Task{
	ID:        "1",
	Title:     "Learn Gin",
	Status:    TaskStatusPending,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var MockTasks = []Task{
	MockTask,
}
