package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (consumer *Consumer) ReceiveMessages(onReceiveMessageHandler func(body []byte)) {
	defer consumer.connection.Close()
	defer consumer.channel.Close()

	msgs, err := consumer.channel.Consume(
		consumer.queue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	failOnError(err, "rabbitmq: failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			onReceiveMessageHandler(d.Body)
		}
	}()

	log.Printf("rabbitmq: Waiting for messages. To exit press CTRL+C")
	<-forever
}

func CreateConsumer(address string) *Consumer {
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
	return &Consumer{connection: conn, channel: ch, queue: &q}

}
