package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLoginForm struct {
	Username string  `json:"username" validate:"required,min=3,max=20"`
	Password *string `json:"password" validate:"required,min=8,max=20"`
}

type UserRegisterForm struct {
	Username string  `json:"username" validate:"required,min=3,max=20"`
	Email    string  `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,min=8,max=20"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"-"`
	Username       *string            `json:"username" bson:"username" validate:"required,min=4,max=20"`
	Email          *string            `json:"email" bson:"email" validate:"required,email"`
	Password       *string            `json:"-" bson:"password" validate:"required,min=8,max=50"`
	CreateDate     time.Time          `json:"-" bson:"createDate"`
	EmailValidated bool               `json:"emailValidated" bson:"emailValidated"`
}

func (user User) String() string {
	return fmt.Sprintf("USER:\nID: %s\n username: %s\n email: %s\n emailVerified: %s\n", user.ID.Hex(), *user.Username, *user.Email, user.EmailValidated)
}


func NewUser(username string, email string, password string) User {
	return User{
		ID:         primitive.NewObjectID(),
		Username:   &username,
		Email:      &email,
		Password:   &password,
		CreateDate: time.Now(),
	}
}
