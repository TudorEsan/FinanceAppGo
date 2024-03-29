package main

import (
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/common"
	"github.com/TudorEsan/FinanceAppGo/server/routes"
	"github.com/hashicorp/go-hclog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var appName = "FinanceAppGo"

func main() {
	// check envs are set

	// dependencies
	router := gin.Default()
	messageClient := common.NewMessagingClient()
	l := hclog.Default()
	l.Info(fmt.Sprintf("Starting %s", appName))
	client := database.DbInstace()

	// react app
	staticServer := static.Serve("/", static.LocalFile("./web", true))
	router.Use(staticServer)

	// cors config
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// routes
	api := router.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	overviewRoutes := api.Group("/overview")
	netWorthRoutes := api.Group("records")
	authRoutes := api.Group("auth")
	base := api.Group("")
	routes.AuthRoutes(authRoutes, l, client, messageClient)
	routes.OverviewRoutes(overviewRoutes, l, client)
	routes.NetWorthRoutes(netWorthRoutes, client, l)
	routes.VerifyRoutes(base, client, l)

	router.NoRoute(func(c *gin.Context) {
		staticServer(c)
	})

	// Start server
	router.Run()
}
