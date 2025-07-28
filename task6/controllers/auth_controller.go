package controllers

import (
	"net/http"
	"task6/services"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := services.RegisterUser(input.UserName, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := services.LoginUser(input.UserName, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

type PromoteInput struct {
	UserName string `json:"user_name" binding:"required"`
}

func Promote(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can promote users"})
		return
	}

	var input PromoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_name is required"})
		return
	}

	err := services.PromoteUser(input.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

