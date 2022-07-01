package helpers

import (
	"App/database"
	"App/models"
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NetWorthCollection *mongo.Collection = database.OpenCollection(database.Client, "NetWorth")
var validate = validator.New()

func InitNetWort(userID string) (netWorth models.NetWorth, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		err = errors.New("could not convert string to primitive")
		return
	}
	records := make([]models.Record, 0)
	netWorth = models.NetWorth{
		Id:      primitive.NewObjectID(),
		UserId:  id,
		Records: records,
	}
	_, err = NetWorthCollection.InsertOne(ctx, netWorth)
	return
}

func AddRecord(userId primitive.ObjectID, record models.Record) (netWorth models.NetWorth, err error) {
	err = validate.Struct(record)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err = NetWorthCollection.FindOneAndUpdate(ctx, bson.M{"userId": userId}, bson.M{
		"$push": bson.M{
			"records": record,
		},
	}, &opts).Decode(&netWorth)
	return
}
