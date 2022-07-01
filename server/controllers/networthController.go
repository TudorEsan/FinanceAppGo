package controller

import (
	"App/database"
	"App/helpers"
	"App/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NetWorth models.NetWorth

var NetWorthCollection *mongo.Collection = database.OpenCollection(database.Client, "NetWorth")

func InitNetWort() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AddRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("HERE?")
		var record models.Record
		if err := c.BindJSON(&record); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		if err := validate.Struct(record); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		record.Id = primitive.NewObjectID()
		newNetWorth, err := helpers.AddRecord(user.ID, record)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"netWorth": newNetWorth,
		})
	}
}
