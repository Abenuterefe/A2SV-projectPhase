package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

// ConnectDB initializes and returns the task collection
func ConnectDB() *mongo.Collection {
	// Replace with your MongoDB URI (use local MongoDB or cloud)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create a new client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Mongo NewClient Error:", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect the client to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Mongo Connect Error:", err)
	}

	// Check if the connection is working
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo Ping Error:", err)
	}

	// Get the collection object
	collection := client.Database("taskdb").Collection("tasks")

	TaskCollection = collection
	return collection
}
