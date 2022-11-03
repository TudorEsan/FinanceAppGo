package controller

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/common"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/helpers"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/models"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/service"
	"github.com/hashicorp/go-hclog"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserController interface {
	StartUpdatingUserAssets()
	StartConsuming()
}

type UserController struct {
	userCollection   *mongo.Collection
	l                hclog.Logger
	rabbitClient     common.IMessagingClient
	assetsCollection *mongo.Collection
}

func NewUserController(l hclog.Logger, mongoClient *mongo.Client, rabbitMq common.IMessagingClient) IUserController {
	l.Named("UserController")
	userCollection := database.OpenCollection(mongoClient, "users")
	assetsCollection := database.OpenCollection(mongoClient, "assets")
	return &UserController{userCollection, l, rabbitMq, assetsCollection}
}

type IdPayload struct {
	Id string `json:"id"`
}

func (c *UserController) createUser(id primitive.ObjectID) error {
	newUser := models.User{
		Id: id,
		BinanceKeys: models.BinanceKeys{
			ApiKey:    "",
			SecretKey: "",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.userCollection.InsertOne(ctx, newUser)
	return err
}

func (c *UserController) deleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.userCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *UserController) handleUserCreated(delivery amqp091.Delivery) {
	var payload IdPayload
	err := json.Unmarshal(delivery.Body, &payload)
	if err != nil {
		c.l.Error("Could not unmarshal payload", "error", err, "payload", string(delivery.Body))
		delivery.Reject(false)
		return
	}
	c.l.Info("Recieved user created", "ID", payload.Id)
	id, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		c.l.Error("Could not parse id", "error", err, "id", payload.Id)
		delivery.Reject(false)
		return
	}
	err = c.createUser(id)
	if err != nil {
		c.l.Error("Could not create user", "error", err, "id", payload.Id)
		delivery.Reject(false)
		return
	}
	delivery.Ack(false)
}

func (c *UserController) handleUserDeleted(delivery amqp091.Delivery) {
	var payload IdPayload
	err := json.Unmarshal(delivery.Body, &payload)
	if err != nil {
		c.l.Error("Could not unmarshal payload", "error", err, "payload", string(delivery.Body))
		return
	}
	c.l.Info("User deleted", "ID", payload.Id)
	id, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		c.l.Error("Could not parse id", "error", err, "id", payload.Id)
		delivery.Reject(false)
		return
	}
	err = c.deleteUser(id)
	if err != nil {
		c.l.Error("Could not delete user", "error", err, "id", payload.Id)
		delivery.Reject(false)
		return
	}

	delivery.Ack(false)
}

func (c *UserController) handleUserAction(delivery amqp091.Delivery) {
	switch delivery.RoutingKey {
	case "user.created":
		c.handleUserCreated(delivery)
	case "user.deleted":
		c.handleUserDeleted(delivery)
	default:
		c.l.Error("Unknown routing key", "routingKey", delivery.RoutingKey)
		delivery.Reject(false)
	}

}

func (c *UserController) ListenForUserCreated() {
	opt := common.SubscribeOpt{
		ExchangeName: "portofolio-server",
		RoutingKeys:  []string{"user.*"},
		QueueName:    "broker-service-user-created",
		HandlerFunc:  c.handleUserAction,
	}
	c.rabbitClient.Subscribe(opt)
}

func (c *UserController) StartConsuming() {
	c.l.Info("Listening for messages")
	go c.ListenForUserCreated()
}

func (c *UserController) StartUpdatingUserAssets() {
	c.l.Info("Starting updating user assets")
	go func() {
		ticker := time.NewTicker(time.Hour * 24)
		for {
			<-ticker.C
			c.l.Info("Updating user assets")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
			defer cancel()
			cursor, err := c.userCollection.Find(ctx, bson.M{})
			if err != nil {
				c.l.Error("Could not get users", "error", err)
				return
			}
			defer cursor.Close(ctx)
			for cursor.Next(ctx) {
				var user models.User
				err := cursor.Decode(&user)
				if err != nil {
					c.l.Error("Could not decode user", "error", err)
					continue
				}
				decryptedSecret := helpers.Decrypt(user.BinanceKeys.SecretKey)
				decryptedApiKey := helpers.Decrypt(user.BinanceKeys.ApiKey)
				c.l.Info("decrypted", "secret", decryptedSecret, "api", decryptedApiKey)
				binanceS := service.NewBinanceService(decryptedApiKey, decryptedSecret)
				assets, err := binanceS.GetAssets()
				if err != nil {
					c.l.Error("Could not get assets", "error", err)
					continue
				}
				// update user assets
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				userAssets := models.UserAssets{
					Assets: assets,
					Date:   time.Now(),
					Id:     primitive.NewObjectID(),
					UserId: user.Id,
				}
				_, err = c.userCollection.InsertOne(ctx, userAssets)
				if err != nil {
					c.l.Error("Could not update user assets", "error", err)
					continue
				}
			}
		}
	}()
}
