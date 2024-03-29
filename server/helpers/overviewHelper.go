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

func GetCurrentYear() int {
	return time.Now().Year()
}

func GetLast2Records(recordCollection *mongo.Collection,userID primitive.ObjectID) (records []models.Record, err error) {
	records = make([]models.Record, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	l.Info("Getting last 2 records", "userID", userID)

	curr, err := recordCollection.Find(ctx, bson.M{"userId": userID}, options.Find().SetSort(bson.M{"date": -1}).SetLimit(2))
	if err != nil {
		return
	}
	for curr.Next(ctx) {
		var auxRecord models.Record
		err = curr.Decode(&auxRecord)
		l.Info("Decoding record", "record", auxRecord)
		if err != nil {
			return
		}
		records = append(records, auxRecord)
	}
	return
}

func GetRecordsOverview(recordCollection *mongo.Collection,userId primitive.ObjectID, limit int) (overview models.Overview, err error) {
	overview.NetworthOverview = make([]models.NetworthOverview, 0)
	overview.LiquidityOverview = make([]models.LiquidityOverview, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	curr, err := recordCollection.Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{
				"userId": bson.M{
					"$eq": userId,
				},
			},
		},
		bson.M{
			"$limit": limit,
		},
		bson.M{
			"$sort": bson.M{
				"date": 1,
			},
		},
		bson.M{
			"$addFields": bson.M{
				"total": bson.M{
					"$add": bson.A{
						"$liquidity",
						"$investedAmount",
					},
				},
			},
		},
	})

	if err != nil {
		return
	}
	for curr.Next(ctx) {
		var netWorthOverview models.NetworthOverview
		var liquidityOverview models.LiquidityOverview
		err = curr.Decode(&netWorthOverview)
		if err != nil {
			return
		}
		err = curr.Decode(&liquidityOverview)
		if err != nil {
			return
		}
		overview.NetworthOverview = append(overview.NetworthOverview, netWorthOverview)
		overview.LiquidityOverview = append(overview.LiquidityOverview, models.LiquidityOverview(liquidityOverview))
	}
	return
}
