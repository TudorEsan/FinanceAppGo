package controller

import (
	"fmt"
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
		var recordBody models.RecordBody
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
		cc.l.Info(fmt.Sprintf("Adding record for user %s", userId))

		record, err := recordBody.ToRecord(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		cc.l.Info(fmt.Sprintf("Adding record: %s", record))

		err = helpers.AddRecord(cc.recordCollection, userId, record)
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
		cc.l.Info(fmt.Sprintf("userId: %s", userId))

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			return
		}
		cc.l.Info(fmt.Sprintf("id: %s", id))

		var recordBody models.RecordBody
		if err := c.BindJSON(&recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		cc.l.Info(fmt.Sprintf("recordBody: %v", recordBody))
		recordBody.Id = id

		if err := validate.Struct(recordBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		record, err := recordBody.ToRecord(userId)
		if err != nil {
			cc.l.Error(fmt.Sprintf("Error converting recordBody to record: %v", err.Error()))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		cc.l.Info(fmt.Sprintf("FROM RecordBody: \n  %v", record))

		updatedRecord, err := helpers.UpdateRecord(cc.recordCollection, userId, record)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		cc.l.Info(fmt.Sprintf("Updated Record: \n  %v", updatedRecord))

		c.JSON(http.StatusOK, updatedRecord)
	}
}
