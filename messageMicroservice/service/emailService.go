package service

import (
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/common"
	"github.com/hashicorp/go-hclog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		hclog.Default().Error(msg, "error", err)
// 		log.Fatal("%s: %s", msg, err)
// 	}
// }

var l = hclog.Default().Named("EmailService")

type IEmailService interface {
	NewEmailService()
	ConnectToBroker()
	Subscribe()
	StartConsuming()
}

type EmailService struct {
	mailClient      *sendgrid.Client
	messagingClient common.IMessagingClient
}

func (e *EmailService) StartConsuming() {

}

func NewEmailService(mailClient *sendgrid.Client) *EmailService {

	msgClient := common.NewMessagingClient()

	return &EmailService{mailClient, msgClient}
}

type EmailOptions struct {
	To          string
	Email       string
	Subject     string
	HtmlContent string
	Content     string
}

func (e *EmailService) sendEmail(opt EmailOptions) error {
	l.Info("Sending email")

	from := mail.NewEmail("Tudor", "tudor.esan@icloud.com")
	subject := opt.Subject
	to := mail.NewEmail(opt.To, opt.Email)

	message := mail.NewSingleEmail(from, subject, to, opt.Content, opt.HtmlContent)
	resp, err := e.mailClient.Send(message)
	l.Info(fmt.Sprintf("Email code: %d", resp.StatusCode))

	return err
}
