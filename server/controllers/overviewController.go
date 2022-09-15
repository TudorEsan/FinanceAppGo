package controller

import (
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNetWorthOverview() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		limit := c.DefaultQuery("limit", "10")
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		overview, err := helpers.GetRecordsOverview(user.ID, limitInt)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		last2Records, err := helpers.GetLast2Records(user.ID)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		overview.CurrentRecord = last2Records[0]
		overview.LastRecord = last2Records[1]
		c.JSON(http.StatusOK, gin.H{"overview": overview})
	}
}
