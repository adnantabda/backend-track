package data

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task_manager/models"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	taskCol    *mongo.Collection
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func InitMongo(uri string, dbName string) error {
	var err error
	ctx, cancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db = client.Database(dbName)
	taskCol = db.Collection("tasks")

	return nil
}

func GetAllTasks() ([]models.Task, error) {
	cursor, err := taskCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	tasks := []models.Task{}
	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid id format")
	}

	var task models.Task
	err = taskCol.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	task.ID = "" // Clear ID to let MongoDB generate it

	res, err := taskCol.InsertOne(context.Background(), task)
	if err != nil {
		return models.Task{}, err
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return models.Task{}, errors.New("failed to convert inserted ID")
	}

	task.ID = objID.Hex()
	return task, nil
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid id format")
	}

	updatedTask.ID = id

	filter := bson.M{"_id": objID}
	replaceRes, err := taskCol.ReplaceOne(context.Background(), filter, updatedTask)
	if err != nil {
		return models.Task{}, err
	}

	if replaceRes.MatchedCount == 0 {
		return models.Task{}, errors.New("task not found")
	}

	return updatedTask, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	deleteRes, err := taskCol.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if deleteRes.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func DisconnectMongo() error {
	cancelFunc()
	return client.Disconnect(context.Background())
}
