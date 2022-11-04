package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/database"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/helpers"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/models"
	"github.com/TudorEsan/FinanceAppGo/BrokerService/service"
	sharedhelpers "github.com/TudorEsan/shared-finance-app-golang/sharedHelpers"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IAssetsController interface {
	GetUserAssets() gin.HandlerFunc
	StartUpdatingUserAssets()
}

type AssetsController struct {
	l                hclog.Logger
	userCollection   *mongo.Collection
	assetsCollection *mongo.Collection
}

func NewAssetsController(l hclog.Logger, mongoClient *mongo.Client) IAssetsController {
	l.Named("AssetsController")
	userCollection := database.OpenCollection(mongoClient, "users")
	assetsCollection := database.OpenCollection(mongoClient, "assets")
	return &AssetsController{l, userCollection, assetsCollection}
}

func (controller *AssetsController) GetUserAssets() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := sharedhelpers.GetUserIdFromCtx(c)
		if err != nil {
			controller.l.Error("Could not get username from context", err)
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		controller.l.Info("Getting user assets", "id", id)
		var page, limit string
		page = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
		p, err := strconv.Atoi(page)
		if err != nil {
			controller.l.Error("Could not convert page to int", err)
			c.JSON(400, gin.H{"message": "page must be a number"})
			return
		}
		l, err := strconv.Atoi(limit)
		if err != nil {
			controller.l.Error("Could not convert limit to int", err)
			c.JSON(400, gin.H{"message": "limit must be a number"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cursor, err := controller.assetsCollection.Find(ctx, bson.M{"userId": id}, options.Find().SetSkip(int64((p-1)*l)).SetLimit(int64(l)))
		if err != nil {
			controller.l.Error("Could not get user assets", err)
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		defer cursor.Close(ctx)
		userAssets := make([]models.UserAssets, 0)
		var userAsset models.UserAssets
		for cursor.Next(ctx) {
			err = cursor.Decode(&userAsset)
			if err != nil {
				controller.l.Error("Could not decode user asset", err)
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			userAssets = append(userAssets, userAsset)
		}
		c.JSON(200, userAssets)
	}
}

func (c *AssetsController) StartUpdatingUserAssets() {
	c.l.Info("Starting updating user assets")
	go func() {
		ticker := time.NewTicker(time.Minute * 1)
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
				decryptedSecret, err := helpers.Decrypt(user.BinanceKeys.SecretKey)
				if err != nil {
					c.l.Error("Could not decrypt secret key", "error", err)
					continue
				}
				decryptedApiKey, err := helpers.Decrypt(user.BinanceKeys.ApiKey)
				if err != nil {
					c.l.Error("Could not decrypt api key", "error", err)
					continue
				}
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
				_, err = c.assetsCollection.InsertOne(ctx, userAssets)
				if err != nil {
					c.l.Error("Could not update user assets", "error", err)
					continue
				}
			}
		}
	}()
}
