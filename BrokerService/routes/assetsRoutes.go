package routes

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/controller"
	middlewares "github.com/TudorEsan/shared-finance-app-golang/sharedMiddlewares"
	"github.com/gin-gonic/gin"
)

func InitAssetRoutes(r *gin.RouterGroup, assetsController controller.IAssetsController) {
	r.Use(middlewares.VerifyAuth())
	r.GET("", assetsController.GetUserAssets())
}
