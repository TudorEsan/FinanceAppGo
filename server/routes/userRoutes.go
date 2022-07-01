package routes

import (
	controller "App/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use(middleware.VerifyAuth())
	incomingRoutes.GET("api/secret", controller.TestMiddlewareController())
}