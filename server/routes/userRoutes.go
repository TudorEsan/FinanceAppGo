package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	// incomingRoutes.Use(middleware.VerifyAuth())
	incomingRoutes.GET("api/secret", controller.TestMiddlewareController())
}