package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Publisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func CreatePublisher(address string) *Publisher {
	conn, err := amqp.Dial(address)
	failOnError(err, "rabbitmq: failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "rabbitmq: failed to open a channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "rabbitmq: failed to declare a queue")

	return &Publisher{connection: conn, channel: ch, queue: &q}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (publisher *Publisher) SendMessage(body string) {
	err := publisher.channel.Publish("", publisher.queue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	failOnError(err, "rabbitmq: failed to publish a message")
}
