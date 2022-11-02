package controller

import (
	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserController interface {
	StartConsuming()
}

type UserController struct {
	mongoCollection *mongo.Collection
	l               hclog.Logger
	rabbitClient    common.IMessagingClient
}

func NewUserController(l hclog.Logger, mongoClient *mongo.Client, rabbitMq common.IMessagingClient) IUserController {
	l.Named("UserController")
	mongoC := database.OpenCollection(mongoClient, "users")
	return &UserController{mongoC, l, rabbitMq}
}

func (c *UserController) handleUserCreated(delivery amqp091.D) {
}

func (c *UserController) ListenForUserCreated() {
	opt := common.SubscribeOpt{
		ExchangeName: "users",
		ExchangeType: "direct",
		RoutingKeys:  []string{"email"},
		QueueOptions: common.QueueOpt{
			Exclusive: true,
		},
		HandlerFunc: e.emailHandlerFunc,
	}
	c.rabbitClient.Consume("user", c.l, c.handleMessage)
}

func (c *UserController) StartConsuming() {
	c.l.Info("Listening for messages")
	go c.ListenForUserCreated()

}
