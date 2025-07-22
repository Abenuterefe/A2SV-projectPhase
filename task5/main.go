package main

import (
	"A2SV-projectPhase/task5/config"
	routes "A2SV-projectPhase/task5/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB and store the collection reference
	config.ConnectDB()

	// Initialize the Gin router
	router := gin.Default()

	// Register task routes
	routes.TaskRoutes(router)

	// Start the server
	router.Run(":8080") // Access the API at http://localhost:8080
}
