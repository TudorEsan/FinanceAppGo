package helpers

import (
	"context"
	"time"

	"github.com/TudorEsan/FinanceAppGo/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddRecord(recordCollection *mongo.Collection, userId primitive.ObjectID, record models.Record) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// add id to the record
	record.Id = primitive.NewObjectID()
	record.UserId = userId
	record.GenerateStatistics()

	_, err = recordCollection.InsertOne(ctx, record)
	return
}

func GetRecords(recordCollection *mongo.Collection, userId primitive.ObjectID, page, limit int) (records []models.Record, err error) {
	records = make([]models.Record, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	l := int64(limit)
	skip := int64(page * limit)
	opt := options.FindOptions{
		Skip:  &skip,
		Limit: &l,
		Sort: bson.M{
			"date": -1,
		},
	}
	curr, err := recordCollection.Find(ctx, bson.M{"userId": userId}, &opt)
	if err != nil {
		return
	}
	for curr.Next(ctx) {
		var record models.Record
		err = curr.Decode(&record)
		if err != nil {
			return
		}
		records = append(records, record)
	}
	return
}

func DeleteRecord(recordCollection *mongo.Collection, userId, recordId primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	recordCollection.FindOneAndDelete(ctx, bson.M{"userId": userId, "_id": recordId})
	return
}

func GetRecord(recordCollection *mongo.Collection, userId primitive.ObjectID, recordId string) (record models.Record, err error) {
	id, err := primitive.ObjectIDFromHex(recordId)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = recordCollection.FindOne(ctx, bson.M{"userId": userId, "_id": id}).Decode(&record)
	return
}

func UpdateRecord(recordCollection *mongo.Collection, userId primitive.ObjectID, record models.Record) (updatedRecord models.Record, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = recordCollection.FindOneAndUpdate(ctx, bson.M{"userId": userId, "_id": record.Id}, bson.M{"$set": record},
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedRecord)
	return updatedRecord, err
}

func GetRecordCount(recordCollection *mongo.Collection, userId primitive.ObjectID) (count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	count, err = recordCollection.CountDocuments(ctx, bson.M{"userId": userId})
	return
}
