package common

import (
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/config"
	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/helpers"
	ampq "github.com/rabbitmq/amqp091-go"
)

type IMessagingClient interface {
	Subscribe(SubscribeOpt)
}

type MessagingClient struct {
	conn *ampq.Connection
}

func NewMessagingClient() *MessagingClient {
	conf := config.New()

	if conf.RABBIT_URL == "" {
		panic("RabbitMQ URL is not set")
	}

	c, err := ampq.Dial(conf.RABBIT_URL)
	if err != nil {
		helpers.FailOnError(err, "Failed to connect to RabbitMQ")
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

func (m *MessagingClient) Subscribe(opt SubscribeOpt) {
	ch, err := m.conn.Channel()
	if err != nil {
		helpers.FailOnError(err, "Failed to open a channel")
	}

	ch.ExchangeDeclare(
		opt.ExchangeName, // name
		opt.ExchangeType, // type
		true,             // durable
		opt.AutoDelete,   // auto-deleted
		opt.Internal,     // internal
		opt.NoWait,       // no-wait
		nil,              // arguments
	)

	q, err := ch.QueueDeclare(
		opt.QueueOptions.QueueName,  // name
		opt.QueueOptions.Durable,    // durable
		opt.QueueOptions.AutoDelete, // delete when unused
		opt.QueueOptions.Exclusive,  // exclusive
		opt.QueueOptions.NoWait,     // no-wait
		nil,                         // arguments
	)

	for _, routingKey := range opt.RoutingKeys {
		err := ch.QueueBind(
			"",               // queue name
			routingKey,       // routing key
			opt.ExchangeName, // exchange
			false,
			nil,
		)
		helpers.FailOnError(err, fmt.Sprintf("Failed to bind a queue: %s, routing: %s", q.Name, routingKey))

	}
	if err != nil {
		helpers.FailOnError(err, "Failed to declare a queue")
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
		helpers.FailOnError(err, "Failed to register a consumer")
	}

	subscribeToMessages(msg, opt.HandlerFunc)

}

func subscribeToMessages(msg <-chan ampq.Delivery, handler func(ampq.Delivery)) {
	for d := range msg {
		handler(d)
	}
}
