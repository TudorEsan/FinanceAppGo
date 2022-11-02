package controller

import (
	"context"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/helpers"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IApiKeyController interface {
	GetBinanceKeys(primitive.ObjectID) (models.BinanceKeys, error)
	SetBinanceKeys() gin.HandlerFunc
}

type ApiKeyController struct {
	config   *config.Config
	l        hclog.Logger
	mongoCol *mongo.Collection
}

func NewApiKeyController(config *config.Config, l hclog.Logger, mongoClient *mongo.Client) IApiKeyController {
	l.Named("ApiKeyController")
	mongoC := database.OpenCollection(mongoClient, "users")
	return &ApiKeyController{config, l, mongoC}
}

// func (controller *ApiKeyController) GetApiKeys() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		username, err := sharedhelpers.GetUserIdFromCtx(c)
// 		if err != nil {
// 			controller.l.Error("Could not get username from context", err)
// 			c.JSON(400, gin.H{"message": err.Error()})
// 			return
// 		}

// 		ctx, cancel := context.WithTimeout(context.Background(), controller.config.MongoTimeout)
// 		defer cancel()
// 		var keys models.BinanceKeys
// 		err = controller.mongoCol.FindOne(ctx, bson.M{"username": username}).Decode(&keys)
// 		if err != nil {
// 			controller.l.Error("Could not find keys", err)
// 			c.JSON(400, gin.H{"message": err.Error()})
// 			return
// 		}
// 	}
// }

func (controller *ApiKeyController) GetBinanceKeys(userId primitive.ObjectID) (models.BinanceKeys, error) {
	ctx, cancel := context.WithTimeout(context.Background(), controller.config.MongoTimeout)
	defer cancel()
	var user models.User
	err := controller.mongoCol.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		controller.l.Error("Could not find keys", err)
		return models.BinanceKeys{}, err
	}
	decryptedSecretKey := helpers.Decrypt(controller.config.EncryptionKey, user.BinanceKeys.SecretKey)
	decryptedApiKey := helpers.Decrypt(controller.config.EncryptionKey, user.BinanceKeys.ApiKey)
	return models.BinanceKeys{
		ApiKey:    decryptedApiKey,
		SecretKey: decryptedSecretKey,
	}, nil
}

func (controller *ApiKeyController) SetBinanceKeys() gin.HandlerFunc {
	return func(c *gin.Context) {
		var keys models.BinanceKeys

		// id, err := sharedhelpers.GetUserIdFromCtx(c)
		idS, err := "63629d3c66553cad887b9963", error(nil)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Could not get username ",
			})
			return
		}

		if err := c.BindJSON(&keys); err != nil {
			controller.l.Error("Could not bind", err)
			c.JSON(400, gin.H{"message": err.Error(), "body": c.Request.Body})
			return
		}

		// encrypt the keys
		encryptedApiKey := helpers.Encrypt([]byte(controller.config.EncryptionKey), keys.ApiKey)
		encryptedSecretKey := helpers.Encrypt([]byte(controller.config.EncryptionKey), keys.SecretKey)
		var user models.User

		ctx, cancel := context.WithTimeout(context.Background(), controller.config.MongoTimeout)
		defer cancel()

		id, err := primitive.ObjectIDFromHex(idS)
		if err != nil {
			controller.l.Error("Could not convert id to object id", err)
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		update := bson.M{
			"$set": bson.M{
				"binanceKeys": bson.M{
					"apiKey":    encryptedApiKey,
					"secretKey": encryptedSecretKey,
				},
			},
		}
		options := options.FindOneAndUpdate().SetReturnDocument(options.After)
		err = controller.mongoCol.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, options).Decode(&user)
		if err != nil {
			controller.l.Error("Could not update keys", err)
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"user": user,
		})
	}
}
