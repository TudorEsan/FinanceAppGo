package helpers

import (
	"App/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	var user models.User
	defer cancel()
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println("valid id")
	err = userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println(user)
	return user, nil
}
