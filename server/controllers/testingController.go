package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func TestMiddlewareController() gin.HandlerFunc {
	return func (c *gin.Context) {
		fmt.Print("hihi in")
		c.JSON(http.StatusOK, gin.H{
			"message": "hehe you are in!",
		})
	}
}