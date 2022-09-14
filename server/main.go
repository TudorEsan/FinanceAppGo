package main

import (
	"App/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// port := os.Getenv("8080")
	// if port == "" {
	// 	port = "8080"
	// }
	router := gin.Default()
	
	// react app
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"*"}

	// cors config
	config.AllowOrigins = []string{"http://localhost:3000"}
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
	routes.AuthRoutes(authRoutes)
	routes.UserRoutes(api)
	routes.OverviewRoutes(overviewRoutes)
	routes.NetWorthRoutes(netWorthRoutes)

	// Start server
	router.Run()
}
