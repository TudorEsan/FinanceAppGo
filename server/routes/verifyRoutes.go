package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func VerifyRoutes(incomingRoutes *gin.RouterGroup, client *mongo.Client, l hclog.Logger) {
	c := controller.NewVerificationController(l, client)
	incomingRoutes.GET("/verify/:verificationToken", c.VerificationHandler())
}
