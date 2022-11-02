package main

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/service"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-hclog"
)

var apiKey = ""
var secretKey = ""

func main() {
	l := hclog.Default()
	mongoClient := database.DbInstace()
	messagingClient := common.NewMessagingClient()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	service.NewBinanceService(apiKey, secretKey, redisClient)
	userController := controller.NewUserController(l, mongoClient, messagingClient)
	userController.StartConsuming()

	never := make(chan bool)
	<- never
}
