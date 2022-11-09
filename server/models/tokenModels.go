package models

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SECRET_KEY []byte = getSecretKey()

type SignedDetails struct {
	Email          string
	Username       string
	Id             string
	EmailValidated bool
	jwt.StandardClaims
}

type EmailVerificationToken struct {
	UserId primitive.ObjectID
	jwt.StandardClaims
}

func getSecretKey() []byte {
	fmt.Println("Getting secret key")
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		fmt.Println("Warning JWT_SECRET not set, using default secret")
		secret = "secret"
	}
	return []byte(secret)
}
