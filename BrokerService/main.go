package main

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/routes"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-hclog"
)

var apiKey = ""
var secretKey = ""

func main() {
	// dependencies
	l := hclog.Default()
	mongoClient := database.DbInstace()
	messagingClient := common.NewMessagingClient()
	config := config.New()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// server
	router := gin.Default()

	// routes
	keysGroup := router.Group("/keys")
	routes.InitKeyRoutes(keysGroup, config, l, mongoClient)

	// listen to messages
	service.NewBinanceService(apiKey, secretKey, redisClient)
	userController := controller.NewUserController(l, mongoClient, messagingClient)
	userController.StartConsuming()

	router.Run()
	never := make(chan bool)
	<-never
}
