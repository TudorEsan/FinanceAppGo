package routes

import (
	controller "App/controllers"
	"App/middleware"

	"github.com/gin-gonic/gin"
)

func NetWorthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.VerifyAuth())
	incomingRoutes.POST("record", controller.AddRecord())
	incomingRoutes.DELETE("record/:id", controller.DeleteRecord())
	incomingRoutes.GET("record/:id", controller.GetRecord())
	incomingRoutes.GET("", controller.GetNetWorth())
}

