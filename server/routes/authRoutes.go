package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup, l hclog.Logger, client *mongo.Client) {
	c := controller.NewAuthController(l, client)

	incomingRoutes.POST("/signup", c.SignupHandler())
	incomingRoutes.POST("/login", c.LoginHandler())
}
