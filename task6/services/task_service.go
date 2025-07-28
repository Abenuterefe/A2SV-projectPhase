package services

import (
	"context"
	"time"
	"errors"
	"task6/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(title, description string, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
		Completed:   false,
		UserID:      userID,
		CreatedAt:   time.Now(),
	}

	_, err := models.TaskCollection.InsertOne(ctx, task)
	return err
}

func GetTasks(userID primitive.ObjectID) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := models.TaskCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}




func UpdateTask(id string, title, description string, completed bool) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	update := bson.M{"$set": bson.M{
		"title":       title,
		"description": description,
		"completed":   completed,
	}}

	_, err = models.TaskCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = models.TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

