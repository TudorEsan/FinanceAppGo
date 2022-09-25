package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	middlewares "github.com/TudorEsan/FinanceAppGo/server/middleware"
	"github.com/hashicorp/go-hclog"

	"github.com/gin-gonic/gin"
)

func OverviewRoutes(incomingRoutes *gin.RouterGroup, l hclog.Logger) {
	c := controller.NewOverviewController(l)
	incomingRoutes.Use(middlewares.VerifyAuth())
	incomingRoutes.GET("/networth", c.GetNetWorthOverview())
}
