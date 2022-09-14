package controller

// func Signup
import (
	"github.com/TudorEsan/FinanceAppGo/server/database"
	helper "github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/TudorEsan/FinanceAppGo/server/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func Login

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// func VerifyUser() gin.Hal{
// 	return
// }
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		fmt.Print("signup")
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
			return
		}
		fmt.Println("befor valid username")
		err := helper.ValidUsername(*user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		user.ID = primitive.NewObjectID()
		fmt.Println("befor hash password")
		hashedPassw, _ := helper.HashPassword(*user.Password)
		*user.Password = hashedPassw
		fmt.Println("before generate tokens")
		jwt, refreshToken, _ := helper.GenerateTokens(user)
		user.RefreshToken = &refreshToken
		fmt.Println("before insert user")
		_, err = userCollection.InsertOne(ctx, user)
		fmt.Println("after insert user")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.SetCookie("token", jwt, 60*60*24*30, "", "", false, false)
		c.SetCookie("refreshToken", jwt, 60*60*24*30, "", "", false, false)
		if err != nil {
			helper.ReturnError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Login")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var user models.User
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("BEFORE FIND ONE")
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, errors.New("not valid username"))
			return
		}
		fmt.Println("BEFORE COMPARE PASSWORD")
		err = helper.CheckPassword(foundUser, user)
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		fmt.Println("BEFORE GENERATE TOKENS")
		jwt, refreshToken, err := helper.GenerateTokens(foundUser)
		if err != nil {
			helper.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		fmt.Println("BEDROE UPDATE REFRESH TOKEN")
		foundUser, err = helper.UpdateTokens(c, jwt, refreshToken, foundUser.ID.Hex())
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		foundUser.Password = nil
		c.JSON(http.StatusOK, foundUser)
	}
}
