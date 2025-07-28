package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "task6/config" 
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName  string             `bson:"user_name" json:"user_name"`
	Password  string             `bson:"password,omitempty" json:"-"`
	Role      string             `bson:"role" json:"role"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var UserCollection *mongo.Collection

func InitUserCollection(db *mongo.Database) {
	UserCollection = db.Collection("users")
}
