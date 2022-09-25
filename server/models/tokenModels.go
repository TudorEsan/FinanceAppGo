package models

import (
	"os"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SECRET_KEY []byte = getSecretKey()

type SignedDetails struct {
	Email    string
	Username string
	Id       string
	jwt.StandardClaims
}

type EmailVerificationToken struct {
	UserId primitive.ObjectID
	jwt.StandardClaims
}

func getSecretKey() []byte {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		panic("SECRET_KEY not found")
	}
	return []byte(secret)
}