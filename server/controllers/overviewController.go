package controller

import (
	"App/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOverview() gin.HandlerFunc {
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
		records, err := helpers.GetYearRecords(user.ID, year)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"records": records})
	}
}
