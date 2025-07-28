package controllers

import (
	"net/http"
	"task6/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTask(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)

	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := services.CreateTask(input.Title, input.Description, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func GetTasks(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)

	tasks, err := services.GetTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}



type UpdateInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}

func UpdateTask(c *gin.Context) {
	var input UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id := c.Param("id")
	err := services.UpdateTask(id, input.Title, input.Description, input.Completed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
