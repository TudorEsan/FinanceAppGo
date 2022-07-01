package main

import (
	"App/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("8080")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	netWorthRoutes := router.Group("networth")
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.NetWorthRoutes(netWorthRoutes)
	router.Run()

}
