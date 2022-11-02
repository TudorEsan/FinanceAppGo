package controller

import (
	"encoding/json"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/hashicorp/go-hclog"
	"github.com/rabbitmq/amqp091-go"
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

type UserCreatedPayload struct {
	Id string `json:"id"`
}

func (c *UserController) handleUserCreated(delivery amqp091.Delivery) {
	var payload UserCreatedPayload
	err := json.Unmarshal(delivery.Body, &payload)
	if err != nil {
		c.l.Error("Could not unmarshal payload", "error", err, "payload", string(delivery.Body))
		return
	}
	c.l.Info("User created", "ID", payload.Id)
	delivery.Ack(false)
}

func (c *UserController) ListenForUserCreated() {
	opt := common.SubscribeOpt{
		ExchangeName: "user",
		ExchangeType: "direct",
		RoutingKeys: []string{"created"},
		QueueName:   "broker-service-user-created2",
		HandlerFunc: c.handleUserCreated,
	}

	c.rabbitClient.Subscribe(opt)
}

func (c *UserController) StartConsuming() {
	c.l.Info("Listening for messages")
	go c.ListenForUserCreated()
}
