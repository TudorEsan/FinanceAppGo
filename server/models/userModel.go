package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Username     *string            `json:"username" bson:"username" validate:"required,min=4,max=20"`
	Email        *string            `json:"email" bson:"email" validate:"required,email"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=8,max=50"`
	RefreshToken *string            `json:"refreshToken" bson:"refreshToken"`
	CreateDate   time.Time          `json:"createDate" bson:"createDate"`
}

func (user User) String() string {
	return fmt.Sprintf("USER:\nID: %s\n username: %s\n email: %s\n refreshToken: %s\n", user.ID.Hex(), *user.Username, *user.Email, *user.RefreshToken)
}
