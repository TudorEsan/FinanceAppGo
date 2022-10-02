package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/TudorEsan/FinanceAppGo/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)



var SECRET_KEY []byte = getSecretKey()

func GenerateTokens(user models.User) (string, string, error) {
	claims := &models.SignedDetails{
		Email:    *user.Email,
		Username: *user.Username,
		Id:       user.ID.Hex(),
		EmailValidated: user.EmailValidated,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 60 * 24 * 30).Unix(),
		},
	}
	refreshClaims := &models.SignedDetails{
		Id: user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24 * 30).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET_KEY)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (*models.SignedDetails, error) {
	token, err := jwt.ParseWithClaims(signedToken, &models.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil && strings.Contains(err.Error(), "expired") {
		return nil, errors.New("token expired")
	}
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func RemoveCookies(c *gin.Context) {
	c.SetCookie("token", "", 60*60*24*30, "", "", false, false)
	c.SetCookie("refreshToken", "", 60*60*24*30, "", "", false, false)
}

func getSecretKey() []byte {
	envs, _ := godotenv.Read(".env")
	return []byte(envs["JWT_SECRET"])
}

func GetUserIdFromVerificationToken(verificationToken string) (primitive.ObjectID, error) {
	token, err := jwt.ParseWithClaims(verificationToken, &models.EmailVerificationToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SECRET_KEY), nil
	})	
	if err != nil {
		return primitive.NilObjectID, err
	}

	claims := token.Claims.(*models.EmailVerificationToken)

	return claims.UserId, nil
}

func GenerateVerificationToken(userId primitive.ObjectID) (string, error) {
	claims := &models.EmailVerificationToken{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 60 * 24 * 30).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(models.SECRET_KEY)
	if err != nil {
		return "", err
	}
	return token, nil
}