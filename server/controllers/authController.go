package controller

// func Signup
import (
	"App/database"
	helper "App/helpers"
	"App/models"
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
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := helper.ValidUsername(*user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.CreateDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		hashedPassw, _ := helper.HashPassword(*user.Password)
		*user.Password = hashedPassw
		jwt, refreshToken, _ := helper.GenerateTokens(user)
		user.RefreshToken = &refreshToken
		_, err = userCollection.InsertOne(ctx, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("token", jwt, 60*60*24*30, "/", "/", true, true)
		fmt.Println(user.ID.Hex())
		netWorth, err := helper.InitNetWort(user.ID.Hex())
		if err != nil {
			helper.ReturnError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":     user,
			"netWorth": netWorth,
		})
	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var user models.User
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Bind JSON :", time.Since(start).Milliseconds())
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		fmt.Println("Find Username", time.Since(start).Milliseconds())
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, errors.New("not valid username"))
			return
		}
		err = helper.CheckPassword(foundUser, user)
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		fmt.Println("Check Password ", time.Since(start).Milliseconds())
		jwt, refreshToken, _ := helper.GenerateTokens(foundUser)
		fmt.Println("Generate Tokens", time.Since(start).Milliseconds())
		fmt.Println(jwt, "       ", refreshToken)
		foundUser, err = helper.UpdateTokens(c, jwt, refreshToken, foundUser.ID.Hex())
		fmt.Println("Generate Tokens ", time.Since(start).Milliseconds())
		if err != nil {
			helper.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}
