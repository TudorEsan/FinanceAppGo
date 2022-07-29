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
var InfoCollection *mongo.Collection = database.OpenCollection(database.Client, "Info")
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

func AddInfo(userId primitive.ObjectID, info models.Info) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = InfoCollection.InsertOne(ctx, info)
	return
}

func GetNetWorth(userId primitive.ObjectID) (netWorth models.NetWorth, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = NetWorthCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&netWorth)
	return
}

func DeleteRecord(userId, recordId primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	record, err := GetRecord(userId, recordId.Hex())
	if err != nil {
		return
	}
	InfoCollection.DeleteOne(ctx, bson.M{"_id": record.InfoId})
	NetWorthCollection.FindOneAndUpdate(ctx, bson.M{"userId": userId},
		bson.M{"$pull": bson.M{
			"records": bson.M{"id": recordId},
		}}, &ReturnNewObject)
	return
}

func GetRecord(userId primitive.ObjectID, recordId string) (record models.Record, err error) {
	var netWorth models.NetWorth
	id, err := primitive.ObjectIDFromHex(recordId)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = NetWorthCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&netWorth)
	// TODO: find how to get subarray from document
	if err != nil {
		return
	}
	found := false
	for _, record = range netWorth.Records {
		if record.Id == id {
			found = true
			break
		}
	}
	if !found {
		err = errors.New("record not found")
		return
	}
	return
}

func GetRecordWithInfo(userId primitive.ObjectID, recordId string) (recordWithInfo models.RecordBody, err error) {
	var info models.Info
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	record, err := GetRecord(userId, recordId)
	if err != nil {
		return
	}
	err = InfoCollection.FindOne(ctx, bson.M{"_id": record.InfoId}).Decode(&info)
	if err != nil {
		return
	}
	recordWithInfo = models.ConcatRecord(record, info)
	return
}
