package main

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/routes"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func main() {
	// dependencies
	l := hclog.Default()
	mongoClient := database.DbInstace()
	messagingClient := common.NewMessagingClient()
	config := config.New()

	// init controllers 
	userController := controller.NewUserController(l, mongoClient, messagingClient)
	userController.StartConsuming()
	userController.StartUpdatingUserAssets()

	// server
	router := gin.Default()

	// routes
	keysGroup := router.Group("/keys")
	routes.InitKeyRoutes(keysGroup, config, l, mongoClient)

	router.Run()
	never := make(chan bool)
	<-never
}
