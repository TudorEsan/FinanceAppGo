package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	BinanceKeys BinanceKeys        `json:"binanceKeys" bson:"binanceKeys"`
}

type UserAssets struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Assets []Asset            `json:"assets" bson:"assets"`
	Date   time.Time          `json:"date" bson:"date"`
}
