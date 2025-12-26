package task

import (
	"context"
	"time"

	"github.com/tiongMax/gintaskic/internal/database"
	"github.com/tiongMax/gintaskic/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTasks retrieves tasks, optionally filtered by status
func GetAllTasks(status string) ([]model.Task, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	// Ensure empty slice instead of nil for JSON
	if tasks == nil {
		tasks = []model.Task{}
	}

	return tasks, nil
}

// GetTaskByID finds a task by its ID
func GetTaskByID(id string) (*model.Task, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err // Invalid ID format
	}

	var task model.Task
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Not found
		}
		return nil, err
	}

	return &task, nil
}

func CreateTask(task model.Task) (*model.Task, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
