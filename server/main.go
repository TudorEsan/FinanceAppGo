package main

import (
	"App/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("8080")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	config := cors.DefaultConfig()

	// cors config
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	// serve react app
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	// routes
	api := router.Group("/api")
	netWorthRoutes := api.Group("records")
	authRoutes := api.Group("auth")
	routes.AuthRoutes(authRoutes)
	routes.UserRoutes(api)
	routes.NetWorthRoutes(netWorthRoutes)

	// Start server
	router.Run()
}
