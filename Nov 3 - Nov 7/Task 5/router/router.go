package router

import (
	"github.com/gin-gonic/gin"
	"Nov 3 - Nov 7/Task 5/controllers"
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
