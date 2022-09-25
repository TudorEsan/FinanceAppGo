package controller

import (
	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificationController struct {
	l      hclog.Logger
	collection *mongo.Collection
}

func NewVerificationController(l hclog.Logger, client *mongo.Client) *VerificationController {
	collection := database.OpenCollection(client, "users")
	return &VerificationController{l, collection}
}

func getUserIdFromVerificationToken(verificationToken string) (primitive.ObjectID, error) {
	token, err := jwt.ParseWithClaims(verificationToken, &models.EmailVerificationToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SECRET_KEY), nil
	})	
	if err != nil {
		return primitive.NilObjectID, err
	}

	claims := token.Claims.(*models.EmailVerificationToken)

	return claims.UserId, nil
}

func (cc *VerificationController) VerificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		verificationToken, ok := c.Params.Get("verificationToken")
		cc.l.Info("Verification token: ", verificationToken)
		if !ok {
			cc.l.Error("Verification token not found")
			c.JSON(401, gin.H{"message": "Verification Token not found"})
			return
		}
		userId, err := getUserIdFromVerificationToken(verificationToken)
		cc.l.Info("User id: ", userId)
		if err != nil {
			cc.l.Error("Could not get user id from verification token")
			c.JSON(400, gin.H{"message": "Invalid Verification Token"})
			return
		}
		err = helpers.VerifyUserEmail(cc.collection, userId)
		if err != nil {
			c.JSON(400, gin.H{"message": "Could not verify user email"})
			return
		}
		c.JSON(200, gin.H{"message": "Email verified"})
	}
}
