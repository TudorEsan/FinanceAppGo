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
		bson.D{{"$match", bson.D{{"date", bson.D{{"$gte", time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)}}}}}},
		bson.M{
			"$group": bson.M{
				"_id":            bson.M{"$month": "$date"},
				"liquidity":      bson.M{"$last": "$liquidity"},
				"investedAmount": bson.M{"$last": "$investedAmount"},
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
				"month": "$_id",
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
