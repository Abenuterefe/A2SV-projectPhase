package routes

import (
	"task6/controllers"
	"task6/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	task := router.Group("/tasks")
	task.Use(middleware.AuthMiddleware()) // All tasks routes require login

	// Admin-only routes
	task.POST("/", middleware.IsAdminMiddleware(), controllers.CreateTask)
	task.PUT("/:id", middleware.IsAdminMiddleware(), controllers.UpdateTask)
	task.DELETE("/:id", middleware.IsAdminMiddleware(), controllers.DeleteTask)

	// Optional: public route to get tasks
	task.GET("/", controllers.GetTasks)
}
