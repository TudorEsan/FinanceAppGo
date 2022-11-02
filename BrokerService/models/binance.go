package models

type BinanceKeys struct {
	ApiKey    string `json:"apiKey" bson:"apiKey" binding:"required"`
	SecretKey string `json:"secretKey" bson:"secretKey" binding:"required"`
}