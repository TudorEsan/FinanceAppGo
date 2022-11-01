package routes

import (
	controller "github.com/TudorEsan/FinanceAppGo/server/controllers"
	"github.com/hashicorp/go-hclog"
	"github.com/TudorEsan/FinanceAppGo/server/common"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup, l hclog.Logger, client *mongo.Client, messagingClient *common.IMessagingClient) {
	c := controller.NewAuthController(l, client, messagingClient)

	incomingRoutes.POST("/signup", c.SignupHandler())
	incomingRoutes.POST("/login", c.LoginHandler())
}
