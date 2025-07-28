package config

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    // Get Mongo URI and DB name from environment
    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("DB_NAME")

    if uri == "" || dbName == "" {
        log.Fatal("❌ MONGO_URI or DB_NAME not set in environment")
    }

    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI(uri)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database(dbName)
    log.Println("✅ Connected to MongoDB at", uri)
}
