package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	"github.com/TudorEsan/FinanceAppGo/server/middleware"

	"github.com/gin-gonic/gin"
)

func OverviewRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.VerifyAuth())
	incomingRoutes.GET("/networth", controller.GetNetWorthOverview())
}