package usecases

import (
	"time"

	"task_manager/Domain"
)

// TaskUseCaseImpl implements the TaskUseCase interface
type TaskUseCaseImpl struct {
	taskRepo domain.TaskRepository
}

// NewTaskUseCase creates a new TaskUseCaseImpl instance
func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCaseImpl {
	return &TaskUseCaseImpl{
		taskRepo: taskRepo,
	}
}

// GetAllTasks retrieves all tasks
func (uc *TaskUseCaseImpl) GetAllTasks() ([]domain.Task, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask retrieves a task by ID
func (uc *TaskUseCaseImpl) GetTask(id string) (domain.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

// CreateTask creates a new task
func (uc *TaskUseCaseImpl) CreateTask(task domain.Task) (domain.Task, error) {
	// Validate required fields
	if task.Title == "" || task.Status == "" {
		return domain.Task{}, domain.ErrInvalidInput
	}

	// Validate due date if provided
	if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
		return domain.Task{}, domain.ErrInvalidDueDate
	}

	// Create the task
	createdTask, err := uc.taskRepo.Create(task)
	if err != nil {
		return domain.Task{}, err
	}

	return createdTask, nil
}

// UpdateTask updates an existing task
func (uc *TaskUseCaseImpl) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	// Check if task exists
	existingTask, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return domain.Task{}, err
	}

	// Update fields
	task.ID = existingTask.ID

	// Update the task
	updatedTask, err := uc.taskRepo.Update(id, task)
	if err != nil {
		return domain.Task{}, err
	}

	return updatedTask, nil
}

// DeleteTask deletes a task by ID
func (uc *TaskUseCaseImpl) DeleteTask(id string) error {
	// Check if task exists
	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Delete the task
	err = uc.taskRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
