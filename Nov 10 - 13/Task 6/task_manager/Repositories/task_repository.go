package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task_manager/Domain"
)

// TaskRepositoryMongo implements the TaskRepository interface for MongoDB
type TaskRepositoryMongo struct {
	collection *mongo.Collection
}

// NewTaskRepositoryMongo creates a new TaskRepositoryMongo instance
func NewTaskRepositoryMongo(db *mongo.Database) *TaskRepositoryMongo {
	return &TaskRepositoryMongo{
		collection: db.Collection("tasks"),
	}
}

// GetAll retrieves all tasks from the database
func (r *TaskRepositoryMongo) GetAll() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetByID retrieves a task by its ID
func (r *TaskRepositoryMongo) GetByID(id string) (domain.Task, error) {
	var task domain.Task
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, domain.ErrTaskNotFound
		}
		return task, err
	}

	return task, nil
}

// Create adds a new task to the database
func (r *TaskRepositoryMongo) Create(task domain.Task) (domain.Task, error) {
	task.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.Background(), task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

// Update modifies an existing task in the database
func (r *TaskRepositoryMongo) Update(id string, task domain.Task) (domain.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = r.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, domain.ErrTaskNotFound
		}
		return domain.Task{}, err
	}

	task.ID = id
	return task, nil
}

// Delete removes a task from the database
func (r *TaskRepositoryMongo) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrTaskNotFound
	}

	return nil
}
