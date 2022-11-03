package routes

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	middlewares "github.com/TudorEsan/shared-finance-app-golang/sharedMiddlewares"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitKeyRoutes(r *gin.RouterGroup, config *config.Config, l hclog.Logger, mongoClient *mongo.Client) {
	keysController := controller.NewApiKeyController(config, l, mongoClient)
	r.Use(middlewares.VerifyAuth())
	r.POST("binanceKeys", keysController.SetBinanceKeys())
}
