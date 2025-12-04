package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController, taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	// Auth routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/:id/promote", middleware.AdminOnly(), authController.PromoteUser)
		}

		// Task routes
		tasks := api.Group("/tasks")
		{
			tasks.GET("", taskController.GetAllTasks)
			tasks.GET("/:id", taskController.GetTask)
			tasks.POST("", taskController.CreateTask)
			tasks.PUT("/:id", taskController.UpdateTask)
			tasks.DELETE("/:id", taskController.DeleteTask)
		}
	}

	return r
}
