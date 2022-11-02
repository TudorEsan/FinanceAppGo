package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	BinanceKeys BinanceKeys        `json:"binanceKeys" bson:"binanceKeys"`
}
