package main

import (
	"log"
	// "os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"task6/config"
	"task6/routes"
	"task6/models"
)

func main() {
	config.ConnectDB()
	models.InitUserCollection(config.DB)
	models.InitTaskCollection(config.DB)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env")
	}

	

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.TaskRoutes(r)


	log.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8081")
}
