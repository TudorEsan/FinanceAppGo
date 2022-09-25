package main

import (
	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/routes"
	"github.com/hashicorp/go-hclog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// dependencies
	router := gin.Default()
	l := hclog.Default()
	client := database.DbInstace()

	// react app
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

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
	routes.AuthRoutes(authRoutes, l, client)
	routes.OverviewRoutes(overviewRoutes, l, client)
	routes.NetWorthRoutes(netWorthRoutes, client, l)

	// Start server
	router.Run()
}
