package controller

// func Signup
import (
	"context"
	"net/http"
	"time"

	"github.com/TudorEsan/FinanceAppGo/server/customValidators"
	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	helper "github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/hashicorp/go-hclog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func Login

var validate = validator.New()

type AuthController struct {
	l              hclog.Logger
	userCollection *mongo.Collection
}

func NewAuthController(l hclog.Logger, client *mongo.Client) *AuthController {
	collection := database.OpenCollection(client, "user")
	ll := l.Named("AuthController")
	return &AuthController{ll, collection}
}

func (controller *AuthController) saveUser(ctx context.Context, user models.User) error {
	_, err := controller.userCollection.InsertOne(ctx, user)
	return err
}

func (controller *AuthController) SignupHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			controller.l.Error("Could not bind", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "body": c.Request.Body})
			return
		}
		// check if username is not present in the database
		err := helper.ValidUsername(ctx, controller.userCollection, *user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// apply logic to the user, hash password, add creation date
		userForDb, err := helper.GetUserForDb(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// insert user in the db
		err = controller.saveUser(ctx, userForDb)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// generate all the auth tokens
		jwt, refreshToken, err := helper.GenerateTokens(user)
		if err != nil {
			controller.l.Error("Could not generate tokens", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		helper.SetCookies(c, jwt, refreshToken)

		// send verification email

		verificationToken, err := helpers.GenerateVerificationToken(userForDb.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate verification token"})
			return
		}
		controller.l.Info("Verification token", verificationToken)

		err = helper.SendVerificationEmail(userForDb, verificationToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not send verification email"})
			return
		}
		controller.l.Info("Verification email sent")

		c.JSON(http.StatusOK, gin.H{
			"user": userForDb,
		})
	}

}

func (controller *AuthController) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var user models.UserLoginForm
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err := controller.userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Username does not exist"})
			return
		}

		err = helper.CheckPassword(foundUser, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		jwt, refreshToken, err := helper.GenerateTokens(foundUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate tokens"})
			return
		}

		helpers.SetCookies(c, jwt, refreshToken)
		c.JSON(http.StatusOK, foundUser)
	}
}

func (controller *AuthController) RefreshTokensHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			c.JSON(401, gin.H{"message": "Refresh Token not present"})
			return
		}

		claims, err := customValidators.ValidateToken(refreshToken)
		if err != nil {
			controller.l.Error("Invalid Refresh Token")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid refresh token"})
			return
		}

		user, err := helper.GetUser(controller.userCollection, claims.Id)
		if err != nil {
			controller.l.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		token, refreshToken, err := helper.GenerateTokens(user)
		if err != nil {
			controller.l.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		helper.SetCookies(c, token, refreshToken)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
