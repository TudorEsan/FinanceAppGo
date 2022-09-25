package middlewares

import (
	"net/http"

	"github.com/TudorEsan/FinanceAppGo/server/customErrors"
	"github.com/TudorEsan/FinanceAppGo/server/customValidators"
	"github.com/gin-gonic/gin"
)

func RemoveCookies(c *gin.Context) {
	c.SetCookie("token", "", 60*60*24*30, "", "", false, false)
	c.SetCookie("refreshToken", "", 60*60*24*30, "", "", false, false)
}

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if token exists
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Not Found"})
			RemoveCookies(c)
			c.Abort()
			return
		}
		// Check if Refresh Token exists
		_, err = c.Cookie("refreshToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh Token Not Found"})
			RemoveCookies(c)
			c.Abort()
			return
		}

		// Validate Token
		claims, err := customValidators.ValidateToken(token)
		switch e := err.(type) {

		case nil:
			// token ok -> user authorized
			c.Set("UserId", claims.Id)
			c.Next()
			return
			
		case *customErrors.ExpiredToken:
			// Token expired -> client should refresh the tokens
			c.JSON(http.StatusInternalServerError, gin.H{"message": "token expired"})
			c.Abort()
			return
		default:
			// Token invalid or any other error-> reject Request
			c.JSON(http.StatusUnauthorized, customErrors.GetJsonError(e))
			c.Abort()
			return
		}
	}
}
