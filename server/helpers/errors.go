package helpers

import (
	"github.com/gin-gonic/gin"
)


func ReturnError(c *gin.Context, code int,err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}