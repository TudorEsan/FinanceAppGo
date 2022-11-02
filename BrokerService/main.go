package main

import (
	"context"
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/service"
	"github.com/go-redis/redis/v8"
)

var apiKey = ""
var secretKey = ""

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}


	binanceS := service.NewBinanceService(apiKey, secretKey, redisClient)
	assets, err := binanceS.GetAssets()
	if err != nil {
		panic(err)
	}
	fmt.Println(assets)
}
