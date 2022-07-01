package middleware

import (
	"App/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId string
		token, err := c.Cookie("token")
		if err != nil {
			helpers.ReturnError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(token)
		if err != nil && err.Error() != "token expired" {
			helpers.ReturnError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		if err != nil && err.Error() == "token expired" {
			refreshToken, err := c.Cookie("refreshToken")
			if err != nil {
				helpers.ReturnError(c, http.StatusUnauthorized, err)
				c.Abort()
				return
			}
			user, err := helpers.ValidateRefreshToken(refreshToken)
			if err != nil {
				helpers.ReturnError(c, http.StatusUnauthorized, err)
				c.Abort()
				return
			}
			fmt.Println("token was expired, but refreshtoken was valid")
			fmt.Println("IDDDD: ", user.ID.Hex())
			userId = user.ID.Hex()
			fmt.Println("trece?")

		} else {
			userId = claims.Id
		}
		fmt.Println("??")
		user, err := helpers.GetUser(userId)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}
		newToken, newRefreshToken, err := helpers.GenerateTokens(user)
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}
		user, err = helpers.UpdateTokens(c, newToken, newRefreshToken, user.ID.Hex())
		if err != nil {
			helpers.ReturnError(c, http.StatusBadRequest, err)
		}
		c.Set("user", user)
		c.Next()
	}
}
