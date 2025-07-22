package routes

import (
	"A2SV-projectPhase/task5/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("", controllers.CreateTask)
		taskGroup.GET("", controllers.GetTasks)
		taskGroup.GET("/:id", controllers.GetTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}
