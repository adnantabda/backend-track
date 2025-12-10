package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task_manager/Domain"
)

// TaskController handles HTTP requests for tasks
type TaskController struct {
	taskUseCase domain.TaskUseCase
}

// NewTaskController creates a new TaskController
func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

// GetTasks handles GET /tasks
func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskUseCase.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// GetTask handles GET /tasks/:id
func (c *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := c.taskUseCase.GetTask(id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		}
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// CreateTask handles POST /tasks
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdTask, err := c.taskUseCase.CreateTask(task)
	if err != nil {
		switch err {
		case domain.ErrInvalidInput:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case domain.ErrInvalidDueDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, createdTask)
}

// UpdateTask handles PUT /tasks/:id
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedTask, err := c.taskUseCase.UpdateTask(id, task)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case domain.ErrInvalidInput:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		}
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles DELETE /tasks/:id
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.taskUseCase.DeleteTask(id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

// UserController handles HTTP requests for users
type UserController struct {
	userUseCase domain.UserUseCase
}

// NewUserController creates a new UserController
func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// Register handles POST /register
func (c *UserController) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdUser, err := c.userUseCase.Register(user)
	if err != nil {
		switch err {
		case domain.ErrInvalidInput:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case domain.ErrEmailAlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// Login handles POST /login
func (c *UserController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := c.userUseCase.Login(loginData.Email, loginData.Password)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// GetProfile handles GET /profile
func (c *UserController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := c.userUseCase.GetUserProfile(userID.(string))
	if err != nil {
		if err == domain.ErrUserNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}
