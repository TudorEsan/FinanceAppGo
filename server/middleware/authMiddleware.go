package middlewares

import (
	"fmt"
	"net/http"
	"github.com/TudorEsan/FinanceAppGo/server/customErrors"
	"github.com/TudorEsan/FinanceAppGo/server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func RemoveCookies(c *gin.Context) {
	c.SetCookie("token", "", 60*60*24*30, "", "", false, false)
	c.SetCookie("refreshToken", "", 60*60*24*30, "", "", false, false)
}

type AuthMiddlewareController struct {
	l hclog.Logger
}

func NewAuthMiddlewareController(l hclog.Logger) *AuthMiddlewareController {
	ll := l.Named("AuthMiddlewareController")
	return &AuthMiddlewareController{ll}
}

func (cc *AuthMiddlewareController) VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if token exists
		token, err := c.Cookie("token")
		cc.l.Info("Token: " + token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Not Found"})
			cc.l.Error("Token Not Found")
			RemoveCookies(c)
			c.Abort()
			return
		}
		// Check if Refresh Token exists
		_, err = c.Cookie("refreshToken")
		cc.l.Info("Refresh Token: " + token)
		if err != nil {
			cc.l.Error("Refresh Token Not Found")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh Token Not Found"})
			RemoveCookies(c)
			c.Abort()
			return
		}

		// Validate Token
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			cc.l.Error("Token Validation Error: " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Validation Error"})
			RemoveCookies(c)
			c.Abort()
			return
		}
		

		cc.l.Info(fmt.Sprintf(
			"Claims: %v",
			claims))
		if !claims.EmailValidated {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Email not validated"})
			RemoveCookies(c)
			c.Abort()
			return
		}

		switch e := err.(type) {
		case nil:
			// token ok -> user authorized
			cc.l.Info("Token Valid")
			c.Set("UserId", claims.Id)
			c.Next()
			return

		case *customErrors.ExpiredToken:
			// Token expired -> client should refresh the tokens
			cc.l.Error("Token Expired")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "token expired"})
			c.Abort()
			return
		default:
			// Token invalid or any other error-> reject Request
			cc.l.Error("Token Invalid")
			c.JSON(http.StatusUnauthorized, customErrors.GetJsonError(e))
			c.Abort()
			return
		}
	}
}
