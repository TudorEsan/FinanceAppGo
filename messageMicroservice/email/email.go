package email

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-hclog"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func failOnError(err error, msg string) {
	if err != nil {
		hclog.Default().Error(msg, "error", err)
		log.Fatal("%s: %s", msg, err)
	}
}

type EmailService struct {
	l          hclog.Logger
	mailClient *sendgrid.Client
	conn       *amqp.Connection
}

func NewEmailService(l hclog.Logger, mailClient *sendgrid.Client, conn *amqp.Connection) *EmailService {
	return &EmailService{l, mailClient, conn}
}

type EmailOptions struct {
	To          string
	Email       string
	Subject     string
	HtmlContent string
	Content     string
}

func (e *EmailService) SendEmail(opt EmailOptions) error {
	e.l.Info("Sending email")

	from := mail.NewEmail("Tudor", "tudor.esan@icloud.com")
	subject := opt.Subject
	to := mail.NewEmail(opt.To, opt.Email)

	message := mail.NewSingleEmail(from, subject, to, opt.Content, opt.HtmlContent)
	resp, err := e.mailClient.Send(message)
	e.l.Info(fmt.Sprintf("Email code: %d", resp.StatusCode))

	return err
}

func (e *EmailService) ListenForEmails() {
	ch, err := e.conn.Channel()
	failOnError(err, "Failed to open email channel")

	emailQueue, err := ch.Consume(
		"sendEmail", // name
		"",          // consumerTag,
		true,        // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)

}
