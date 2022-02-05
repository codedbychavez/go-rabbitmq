package main

import (
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ Server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") 
	// conn, err := amqp.Dial("amqp://guest:rabbitmq")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create RabbitMQ channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create the Queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false, 	// durable
		false, 	// delete when unused
		false, 	// eclusive
		false, 	// no-wait
		nil, 	// arguments
	)

	failOnError(err, "failed to declare a queue")

	body := "hello World!"
	err = ch.Publish(
		"", 	// exhange
		q.Name, // routing key
		false, 	// mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		failOnError(err, "Failed to publish a message")
}