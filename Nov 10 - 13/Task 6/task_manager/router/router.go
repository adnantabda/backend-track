package router

import (
	"github.com/gin-gonic/gin"
	"Nov 10 - 13/Task 6/task_manager/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.GET("", controllers.GetTasks)
		tasks.GET(":id", controllers.GetTask)
		tasks.POST("", controllers.CreateTask)
		tasks.PUT(":id", controllers.UpdateTask)
		tasks.DELETE(":id", controllers.DeleteTask)
	}

	return r
}
