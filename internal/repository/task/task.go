package task

import (
	"context"
	"time"

	"github.com/tiongMax/gintaskic/internal/database"
	"github.com/tiongMax/gintaskic/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllTasks retrieves tasks with search, pagination and status filtering
func GetAllTasks(status string, search string, page int64, limit int64) ([]model.Task, int64, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	// Status Filter
	if status != "" {
		filter["status"] = status
	}

	// Search Filter (Case-insensitive regex on title)
	if search != "" {
		filter["title"] = bson.M{"$regex": search, "$options": "i"}
	}

	// Calculate pagination options
	findOptions := options.Find()
	if page > 0 && limit > 0 {
		skip := (page - 1) * limit
		findOptions.SetSkip(skip)
		findOptions.SetLimit(limit)
	}
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}}) // Newest first

	// Get total count for pagination metadata
	totalCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, 0, err
	}

	// Ensure empty slice instead of nil for JSON
	if tasks == nil {
		tasks = []model.Task{}
	}

	return tasks, totalCount, nil
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

// UpdateTask updates an existing task with a map of fields (supports partial updates)
func UpdateTask(id string, updateFields map[string]interface{}) (*model.Task, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Always update updated_at
	updateFields["updated_at"] = time.Now()

	update := bson.M{
		"$set": updateFields,
	}

	var updatedTask model.Task
	// Return the updated document
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts).Decode(&updatedTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Not found
		}
		return nil, err
	}

	return &updatedTask, nil
}

// DeleteTask removes a task by ID
func DeleteTask(id string) error {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// DeleteTasksByStatus removes all tasks with a specific status
func DeleteTasksByStatus(status string) (int64, error) {
	collection := database.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Safety check: require a status
	if status == "" {
		return 0, nil
	}

	result, err := collection.DeleteMany(ctx, bson.M{"status": status})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
