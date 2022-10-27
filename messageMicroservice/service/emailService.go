package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/common"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
	ampq "github.com/rabbitmq/amqp091-go"
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
	l.Info("Started Consuming...")
	opt := common.SubscribeOpt{
		ExchangeName: "messages",
		ExchangeType: "direct",
		RoutingKeys:  []string{"email"},
		QueueOptions: common.QueueOpt{
			Exclusive: true,
		},
		HandlerFunc: e.emailHandlerFunc,
	}
	e.messagingClient.Subscribe(opt)
}

func NewEmailService(mailClient *sendgrid.Client) *EmailService {

	msgClient := common.NewMessagingClient()

	return &EmailService{mailClient, msgClient}
}

type EmailOptions struct {
	To          string `json:"to" validator:"required"`
	Email       string `json:"email" validator:"requiredemail"`
	Subject     string `json:"subject" validator:"required"`
	HtmlContent string `json:"htmlContent"`
	Content     string `json:"content"`
}

func (opt EmailOptions) String() string {
	return fmt.Sprintf(`To: %s

Email: %s

Subject: %s

HtmlContent: %s

Content: %s

`, opt.To, opt.Email, opt.Subject, opt.HtmlContent, opt.Content)
}

func (e *EmailService) emailHandlerFunc(delivery ampq.Delivery) {
	l.Info("Received a message from the broker:")
	validate := validator.New()
	var opt EmailOptions
	err := json.Unmarshal(delivery.Body, &opt)

	if err != nil {
		l.Error("Failed to unmarshal", "error", err)
		delivery.Reject(true)
		time.Sleep(1 * time.Second)
		return
	}
	l.Info(fmt.Sprintf("Body: %s", opt))

	if err := validate.Struct(opt); err != nil {
		l.Error("Failed to validate", "error", err)
		delivery.Reject(true)
		time.Sleep(1 * time.Second)
		return
	}

	if err := e.sendEmail(opt); err != nil {
		l.Error("Failed to send email", "error", err)
		delivery.Reject(true)
		time.Sleep(1 * time.Second)
		return
	}

	delivery.Ack(false)
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
