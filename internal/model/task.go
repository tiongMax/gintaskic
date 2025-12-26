package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TaskStatusPending    = "pending"
	TaskStatusCompleted  = "completed"
	TaskStatusInProgress = "in_progress"
	TaskStatusCancelled  = "cancelled"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title" binding:"required,min=3"`
	Status    string             `bson:"status" json:"status" binding:"required,oneof=pending completed in_progress cancelled"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
