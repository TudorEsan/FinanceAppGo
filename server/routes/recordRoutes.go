package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	middlewares "github.com/TudorEsan/FinanceAppGo/server/middleware"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func NetWorthRoutes(incomingRoutes *gin.RouterGroup, client *mongo.Client, l hclog.Logger) {
	c := controller.NewRecordController(l, client)
	authMiddleware := middlewares.NewAuthMiddlewareController(l)
	incomingRoutes.Use(authMiddleware.VerifyAuth())
	incomingRoutes.POST("", c.AddRecord())
	incomingRoutes.DELETE(":id", c.DeleteRecord())
	incomingRoutes.GET(":id", c.GetRecord())
	incomingRoutes.PUT(":id", c.UpdateRecord())
	incomingRoutes.GET("", c.GetRecords())
	incomingRoutes.GET("count", c.GetRecordCount())
}
