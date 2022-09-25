package controller

import (
	"net/http"
	"strconv"

	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type OverviewController struct {
	l hclog.Logger
	recordCollection *mongo.Collection
}

func NewOverviewController(l hclog.Logger, client *mongo.Client) *OverviewController {
	ll := l.Named("OverviewController")
	recordCollection := client.Database("financeapp").Collection("records")
	return &OverviewController{ll, recordCollection}
}

func (cc *OverviewController) GetNetWorthOverview() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		limit := c.DefaultQuery("limit", "10")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		overview, err := helpers.GetRecordsOverview(cc.recordCollection, user.ID, limitInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		last2Records, err := helpers.GetLast2Records(cc.recordCollection ,user.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		overview.CurrentRecord = last2Records[0]
		overview.LastRecord = last2Records[1]
		c.JSON(http.StatusOK, gin.H{"overview": overview})
	}
}
