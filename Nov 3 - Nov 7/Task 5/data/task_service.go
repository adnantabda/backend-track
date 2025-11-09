package data

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"Nov 3 - Nov 7/Task 5/models"
)

var (
	tasks = make(map[string]models.Task)
	mutex = &sync.Mutex{}
)

func GetAllTasks() []models.Task {
	mutex.Lock()
	defer mutex.Unlock()

	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func GetTaskByID(id string) (models.Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func CreateTask(task models.Task) models.Task {
	mutex.Lock()
	defer mutex.Unlock()

	task.ID = uuid.New().String()
	tasks[task.ID] = task
	return task
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	updatedTask.ID = id
	tasks[id] = updatedTask
	return updatedTask, nil
}

func DeleteTask(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
