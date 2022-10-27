package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TudorEsan/FinanceAppGo/server/database"
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type OverviewController struct {
	l                hclog.Logger
	recordCollection *mongo.Collection
}

func NewOverviewController(l hclog.Logger, client *mongo.Client) *OverviewController {
	ll := l.Named("OverviewController")
	recordCollection := database.OpenCollection(client, "records")
	return &OverviewController{ll, recordCollection}
}

func (cc *OverviewController) GetNetWorthOverview() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := helpers.GetUserIdFromContext(c)
		limit := c.DefaultQuery("limit", "10")
		if err != nil {
			cc.l.Error(fmt.Sprintf("Could not get user from context: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			cc.l.Error(fmt.Sprintf("Could not convert limit to int: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		overview, err := helpers.GetRecordsOverview(cc.recordCollection, userId, limitInt)
		if err != nil {
			cc.l.Error(fmt.Sprintf("Could not get records overview: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		last2Records, err := helpers.GetLast2Records(cc.recordCollection, userId)
		if err != nil {
			cc.l.Error(fmt.Sprintf("Could not get last 2 records: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		cc.l.Info("last 2 records", last2Records)
		if len(last2Records) == 2 {
			overview.CurrentRecord = &last2Records[0]
			overview.LastRecord = &last2Records[1]
		} else {
			overview.CurrentRecord = nil
			overview.LastRecord = nil
		}
		c.JSON(http.StatusOK, gin.H{"overview": overview})
	}
}
