package data

import (
	"task_manager/models"

	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.db.Create(task).Error
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	result := s.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := s.db.Find(&tasks)
	return tasks, result.Error
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.db.Save(task).Error
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.db.Delete(&models.Task{}, id).Error
}
