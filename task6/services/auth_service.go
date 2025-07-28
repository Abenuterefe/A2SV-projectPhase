package services

import (
	"errors"
	"time"

	"task6/models"
	"task6/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
)

func RegisterUser(UserName,password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existing models.User
	err := models.UserCollection.FindOne(ctx, bson.M{"user_names": UserName}).Decode(&existing)
	if err == nil {
		return errors.New("user already exists")
	}

	hashed, _ := utils.HashPassword(password)
	user := models.User{
		ID:        primitive.NewObjectID(),
		UserName:  UserName,
		Password:  hashed,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	_, err = models.UserCollection.InsertOne(ctx, user)
	return err
}

func LoginUser(UserName, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := models.UserCollection.FindOne(ctx, bson.M{"user_name": UserName}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	return utils.GenerateJWT(user.ID, user.Role)
}
func PromoteUser(userName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_name": userName}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := models.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}