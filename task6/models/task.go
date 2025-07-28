package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
	UserID      primitive.ObjectID `bson:"user_id"`
	CreatedAt   time.Time          `bson:"created_at"`
}

var TaskCollection *mongo.Collection

// Call this in main.go after DB is connected
func InitTaskCollection(db *mongo.Database) {
	TaskCollection = db.Collection("tasks")
}
