package controller

import (
	"net/http"
	"strconv"

	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/hashicorp/go-hclog"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordController struct {
	l                hclog.Logger
	recordCollection *mongo.Collection
}

func NewRecordController(l hclog.Logger, client *mongo.Client) *RecordController {
	recordCollection := database.OpenCollection(client, "records")
	return &RecordController{l, recordCollection}
}

func (cc *RecordController) GetRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userId, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		record, err := helpers.GetRecord(cc.recordCollection, userId, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"record": record})

	}
}

func (cc *RecordController) AddRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var recordBody models.Record
		if err := c.BindJSON(&recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := validate.Struct(recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		userId, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = helpers.AddRecord(cc.recordCollection, userId, recordBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func (cc *RecordController) GetRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		perPage, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		records, err := helpers.GetRecords(cc.recordCollection, userId, int(page), perPage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"records": records})

	}
}

func (cc *RecordController) GetRecordCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		count, err := helpers.GetRecordCount(cc.recordCollection, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"recordCount": count})
	}
}

func (cc *RecordController) DeleteRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			return
		}

		netWorthId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		err = helpers.DeleteRecord(cc.recordCollection, userID, netWorthId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	}
}

func (cc *RecordController) UpdateRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := helpers.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			return
		}
		var recordBody models.Record
		if err := c.BindJSON(&recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := validate.Struct(recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = helpers.UpdateRecord(cc.recordCollection, userId, recordBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Updated"})
	}
}
