package helpers

import (
	"App/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCurrentYear() int {
	return time.Now().Year()
}

func GetYearRecords(userId primitive.ObjectID, year int) (monthsOverview []models.MonthOverview, err error) {
	monthsOverview = make([]models.MonthOverview, 0)
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
			"$group": bson.M{
				"_id":            bson.M{"$month": "$date"},
				"liquidity":      bson.M{"$last": "$liquidity"},
				"investedAmount": bson.M{"$last": "$investedAmount"},
				"date":           bson.M{"$last": "$date"},
			},
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
		var monthOverview models.MonthOverview
		err = curr.Decode(&monthOverview)
		if err != nil {
			return
		}
		monthsOverview = append(monthsOverview, monthOverview)
	}
	return
}
