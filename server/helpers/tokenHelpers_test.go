package helpers

import (
	"fmt"
	"testing"
	"time"

	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/go-playground/assert"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

var l = hclog.Default().Named("test")

func TestGenerateTokens(t *testing.T) {
	user := models.NewUser("test", "test", "test")
	token, refToken, err := GenerateTokens(user)
	assert.NotEqual(t, token, "")
	assert.NotEqual(t, refToken, "")
	assert.Equal(t, err, nil)
}

func TestValidateToken(t *testing.T) {
	user := models.NewUser("test", "test", "test")
	l.Info(fmt.Sprintf("user: %v", *user.Username))
	token, refToken, err := GenerateTokens(user)
	assert.Equal(t, err, nil)
	claims, err := ValidateToken(token)
	assert.Equal(t, err, nil)
	assert.Equal(t, claims.Username, *user.Username)
	claims, err = ValidateToken(refToken)
	assert.Equal(t, err, nil)
	l.Info("username", claims.Username)
	l.Info("id", *user.Username)

	assert.Equal(t, claims.Username, *user.Username)
}

func TestValidateTokenInvalid(t *testing.T) {
	_, err := ValidateToken("invalid token")
	assert.NotEqual(t, err, nil)
}

func GenerateExpiredToken() (string, error) {
	user := models.NewUser("test", "test", "test")
	claims := &models.SignedDetails{
		Email:          *user.Email,
		Username:       *user.Username,
		Id:             user.ID.Hex(),
		EmailValidated: user.EmailValidated,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Local().Add(time.Minute * -2).Unix()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func TestValidateTokenExpired(t *testing.T) {
	user := models.NewUser("test", "test", "test")
	token, _, err := GenerateTokens(user)
	assert.Equal(t, err, nil)
	claims, err := ValidateToken(token)
	assert.Equal(t, err, nil)
	assert.Equal(t, claims.Username, user.Username)
}

func TestTokenClaims(t *testing.T) {
	user := models.NewUser("test", "test", "test")
	token, _, err := GenerateTokens(user)
	assert.Equal(t, err, nil)
	claims, err := ValidateToken(token)
	assert.Equal(t, err, nil)
	assert.Equal(t, claims.Username, user.Username)
	assert.Equal(t, claims.Id, user.ID.Hex())
	assert.Equal(t, claims.Email, user.Email)
	assert.Equal(t, claims.EmailValidated, user.EmailValidated)

}