package main

import (
	"log"

	"github.com/streadway/amqp"
)

type MessageQueue struct {
	name string
	queue *amqp.Queue
	channel *amqp.Channel
	conn *amqp.Connection
}

// Here we set the way error messages are displayed in the terminal.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func New(name string) *MessageQueue {
	// connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create a Queue to send messages to.
	q, err := ch.QueueDeclare(
		name,
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")
	mq := MessageQueue{
		name: name,
		queue: &q,
		channel: ch,
		conn: conn,
	}
	return &mq
}

func (mq MessageQueue) Publish(exchange, route, body string) error {
	err := mq.channel.Publish(
		exchange,
		route,
		false,  // mandatory
		false,  // immediate

		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
		
	failOnError(err, "Failed to publish a message")
	log.Printf("  [x] Enqueued: %s", body)

	return err
}