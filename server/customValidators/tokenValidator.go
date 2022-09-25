package customValidators

import (
	"strings"

	"github.com/TudorEsan/FinanceAppGo/server/customErrors"
	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(signedToken string) (*models.SignedDetails, error) {

	token, err := jwt.ParseWithClaims(signedToken, &models.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SECRET_KEY), nil
	})
	if err != nil && strings.Contains(err.Error(), "expired") {
		return nil, customErrors.ExpiredToken{}
	}
	if err != nil {
		return nil, customErrors.InvalidToken{}
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		return nil, customErrors.InvalidToken{}
	}
	return claims, nil
}
