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
var InfoCollection *mongo.Collection = database.OpenCollection(database.Client, "Info")

func InitNetWort() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		fmt.Println("ID ", id)
		record, err := helpers.GetRecord(user.ID, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"record": record})

	}
}

func AddRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var recordBody models.Record
		if err := c.BindJSON(&recordBody); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		if err := validate.Struct(recordBody); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		newNetWorth, err := helpers.AddRecord(user.ID, recordBody)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"netWorth": newNetWorth,
		})
	}
}

func GetNetWorth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		netWorth, err := helpers.GetNetWorth(user.ID)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"netWorth": netWorth})

	}
}

func DeleteRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		var body models.DeleteRecordBody
		if err = c.BindJSON(&body); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		if err := validate.Struct(body); err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
			return
		}
		id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
		}
		err = helpers.DeleteRecord(user.ID, id)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	}

}
