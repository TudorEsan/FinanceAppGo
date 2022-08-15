package helpers

import (
	"App/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCurrentYear() int {
	return time.Now().Year()
}

func GetLast2Records(userID primitive.ObjectID) (records []models.Record, err error) {
	records = make([]models.Record, 0, 2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	curr, err := RecordCollection.Find(ctx, bson.M{"userId": userID}, options.Find().SetSort(bson.M{"date": -1}).SetLimit(2))
	if err != nil {
		return
	}
	for curr.Next(ctx) {
		var auxRecord models.Record
		err = curr.Decode(&auxRecord)
		if err != nil {
			return
		}
		records = append(records, auxRecord)
	}
	return
}

func GetRecordsOverview(userId primitive.ObjectID, year int) (overview models.Overview, err error) {
	overview.NetworthOverview = make([]models.NetworthOverview, 0)
	overview.LiquidityOverview = make([]models.LiquidityOverview, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	curr, err := RecordCollection.Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{
				"date": bson.M{
					"$gte": time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				"userId": bson.M{
					"$eq": userId,
				},
			},
		},
		bson.M{
			"$limit": 10,
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
