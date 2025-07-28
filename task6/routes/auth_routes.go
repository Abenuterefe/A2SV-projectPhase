package routes

import (
	"github.com/gin-gonic/gin"
	"task6/controllers"
	"task6/middlewares"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		
		auth.POST("/promote", middleware.AuthMiddleware(), controllers.Promote)
	}
}
