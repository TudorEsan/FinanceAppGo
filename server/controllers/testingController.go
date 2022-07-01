package controller

import (
	"App/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddlewareController() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserFromContext(c)
		if err != nil {
			helpers.ReturnError(c, http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "hehe you are in!",
			"user":    user,
		})
	}
}
