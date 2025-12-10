package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Domain"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the application routes
func SetupRouter(
	taskController *controllers.TaskController,
	userController *controllers.UserController,
	jwtService *infrastructure.JWTService,
) *gin.Engine {
	r := gin.Default()

	// Public routes
	api := r.Group("/api")
	{
		// User routes
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/register", userController.Register)
			userRoutes.POST("/login", userController.Login)

			// Protected routes
			authorized := userRoutes.Group("")
			authorized.Use(infrastructure.AuthMiddleware(jwtService))
			{
				authorized.GET("/profile", userController.GetProfile)
			}
		}

		// Task routes
		taskRoutes := api.Group("/tasks")
		taskRoutes.Use(infrastructure.AuthMiddleware(jwtService))
		{
			taskRoutes.GET("", taskController.GetTasks)
			taskRoutes.POST("", taskController.CreateTask)
			taskRoutes.GET("/:id", taskController.GetTask)
			taskRoutes.PUT("/:id", taskController.UpdateTask)
			taskRoutes.DELETE("/:id", taskController.DeleteTask)
		}
	}

	return r
}
