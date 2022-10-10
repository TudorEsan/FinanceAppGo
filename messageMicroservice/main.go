package main

import (
	"log"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/config"
	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/email"
	"github.com/hashicorp/go-hclog"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sendgrid/sendgrid-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		hclog.Default().Error(msg, "error", err)
		log.Fatal("%s: %s", msg, err)
	}
}

func main() {
	l := hclog.Default()
	l.Info("Starting message microservice")
	conf := config.New()

	emailClient := sendgrid.NewSendClient(conf.SENDGRID_API_KEY)
	emailService := email.NewEmailService(l, emailClient, )

	conn, err := amqp.Dial(conf.RABBIT_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	emailQueue, err := ch.Consume(
		"sendEmail", // name
		true,   // durable
		false,   // delete when unused
		true,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	

}
