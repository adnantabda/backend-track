package domain

import (
	"errors"
	"time"
)

// Common errors
var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidDueDate     = errors.New("due date cannot be in the past")
)

// Task represents the core business entity for tasks
type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}

// User represents the core business entity for users
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
}

// TaskRepository defines the interface for task data operations
type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (Task, error)
	Create(task Task) (Task, error)
	Update(id string, task Task) (Task, error)
	Delete(id string) error
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user User) (User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}

// TaskUseCase defines the business logic for task operations
type TaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	DeleteTask(id string) error
}

// UserUseCase defines the business logic for user operations
type UserUseCase interface {
	Register(user User) (User, error)
	Login(email, password string) (string, error)
	GetUserProfile(id string) (User, error)
}
