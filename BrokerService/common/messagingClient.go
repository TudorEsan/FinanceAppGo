package common

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/config"
	"github.com/hashicorp/go-hclog"
	ampq "github.com/rabbitmq/amqp091-go"
)

var l = hclog.Default().Named("MessagingClient")

type IMessagingClient interface {
	Subscribe(SubscribeOpt)
	Publish(exchangeName, routingKey string, body []byte) error
}

type MessagingClient struct {
	conn *ampq.Connection
}

func failOnError(err error, msg string) {
	if err != nil {
		l.Error(msg, "error", err)
	}
}

func NewMessagingClient() *MessagingClient {
	conf := config.New()

if conf.RabbitUrl == "" {
		panic("RabbitMQ URL is not set")
	}

	c, err := ampq.Dial(conf.RabbitUrl)
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	return &MessagingClient{c}
}

type SubscribeOpt struct {
	ExchangeName string
	ExchangeType string
	RoutingKeys  []string
	Internal     bool
	NoWait       bool
	AutoDelete   bool
	QueueOptions QueueOpt
	HandlerFunc  func(ampq.Delivery)
}
type QueueOpt struct {
	QueueName  string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

func (client *MessagingClient) Publish(exchangeName, routingKey string, body any) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ch, err := client.conn.Channel()
	if err != nil {

	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare an exchange: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(
		ctx,
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		ampq.Publishing{
			ContentType: "application/json",
			Body:        json,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}
	return nil
}

func (m *MessagingClient) Subscribe(opt SubscribeOpt) {
	ch, err := m.conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}

	ch.ExchangeDeclare(
		opt.ExchangeName, // name
		opt.ExchangeType, // type
		true,             // durable
		opt.AutoDelete,   // auto-deleted
		opt.Internal,     // internal
		opt.NoWait,       // no-wait
		ampq.Table{
			"x-message-ttl": 60_000,
		},
	)

	q, err := ch.QueueDeclare(
		opt.QueueOptions.QueueName,  // name
		opt.QueueOptions.Durable,    // durable
		opt.QueueOptions.AutoDelete, // delete when unused
		opt.QueueOptions.Exclusive,  // exclusive
		opt.QueueOptions.NoWait,     // no-wait
		ampq.Table{
			"x-message-ttl": 60_000,
		},
	)

	for _, routingKey := range opt.RoutingKeys {
		err := ch.QueueBind(
			"",               // queue name
			routingKey,       // routing key
			opt.ExchangeName, // exchange
			false,
			nil,
		)
		failOnError(err, fmt.Sprintf("Failed to bind a queue: %s, routing: %s", q.Name, routingKey))

	}
	if err != nil {
		failOnError(err, "Failed to declare a queue")
	}

	msg, err := ch.Consume(
		q.Name, // name,
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}

	go subscribeToMessages(msg, opt.HandlerFunc)
	l.Info(fmt.Sprintf("Subscribed to %s", opt.ExchangeName))

}

func subscribeToMessages(msg <-chan ampq.Delivery, handler func(ampq.Delivery)) {
	for d := range msg {
		handler(d)
	}
}
