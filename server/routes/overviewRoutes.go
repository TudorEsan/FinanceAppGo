package routes

import (
	controller "App/controllers"
	"App/middleware"

	"github.com/gin-gonic/gin"
)

func OverviewRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.VerifyAuth())
	incomingRoutes.GET("/networth", controller.GetNetWorthOverview())
}