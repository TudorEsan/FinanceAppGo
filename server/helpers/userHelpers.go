package helpers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TudorEsan/FinanceAppGo/server/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserFromContext(c *gin.Context) (user models.User, err error) {
	userAny, exists := c.Get("user")
	if !exists {
		err = errors.New("key does not exist in context")
		return
	}
	user, ok := userAny.(models.User)
	if !ok {
		err = errors.New("could not convert to user")
		return
	}
	return
}

func GetUserForDb(u models.UserRegisterForm) (user models.User, err error) {
	// formats user to be passed to the db
	hashedPassw, err := HashPassword(*u.Password)
	if err != nil {
		return
	}
	return models.NewUser(u.Username, u.Email, hashedPassw), nil
}

func GetUser(userCollection *mongo.Collection, id string) (user models.User, err error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("could not find user in the db")
	}
	return
}

func VerifyUserEmail(userCollection *mongo.Collection, id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var user models.User
	err := userCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"emailValidated": true}}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func SendVerificationEmail(user models.User, verificationToken string) error {
	from := mail.NewEmail("Tudor", "tudor.esan@icloud.com")
	subject := "Verify your email"
	to := mail.NewEmail(*user.Username, *user.Email)
	hclog.L().Info("DOMAIN: ", os.Getenv("DOMAIN_NAME"))
	content := fmt.Sprintf(`
		<html>

		<body>
			<h1>
			Please verify your email</h1>
			<p>Click <a href='%s/api/verify/%s'>here</a> to verify your email</p>
		</body>

		</html>
	`, os.Getenv("DOMAIN_NAME"), verificationToken)

	message := mail.NewSingleEmail(from, subject, to, "", content)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)
	hclog.L().Info("Email code", resp.StatusCode)
	if err != nil {
		return err
	}
	hclog.L().Info("email sent")
	return nil
}

func Sanitize(str string) string {
	return strings.Trim(strings.ToLower(str), " ")
}
