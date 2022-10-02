package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	middlewares "github.com/TudorEsan/FinanceAppGo/server/middleware"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func OverviewRoutes(incomingRoutes *gin.RouterGroup, l hclog.Logger, client *mongo.Client) {
	c := controller.NewOverviewController(l, client)
	authMiddlewareController := middlewares.NewAuthMiddlewareController(l)
	incomingRoutes.Use(authMiddlewareController.VerifyAuth())
	incomingRoutes.GET("/networth", c.GetNetWorthOverview())
}
