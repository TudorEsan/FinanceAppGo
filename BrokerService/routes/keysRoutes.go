package routes

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitKeyRoutes(r *gin.RouterGroup, config *config.Config, l hclog.Logger, mongoClient *mongo.Client) {
	keysController := controller.NewApiKeyController(config, l, mongoClient)
	r.POST("binanceKeys", keysController.SetBinanceKeys())
}
