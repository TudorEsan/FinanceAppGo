package controller

import (
	"App/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNetWorthOverview() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		year, err := strconv.Atoi(c.DefaultQuery("year", fmt.Sprint(helpers.GetCurrentYear())))
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		overview, err := helpers.GetRecordsOverview(user.ID, year)
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
