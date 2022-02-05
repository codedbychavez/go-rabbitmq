package main

import (
	"log"
	"os"
	"strings"

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
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create RabbitMQ channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create the Queue
	q, err := ch.QueueDeclare(
		"task_queue", 	// name
		true, 			// durable
		false, 			// delete when unused
		false, 			// eclusive
		false, 			// no-wait
		nil, 			// arguments
	)

	failOnError(err, "failed to declare a queue")

	body := bodyForm(os.Args)
	err = ch.Publish(
		"", 	// exhange
		q.Name, // routing key
		false, 	// mandatory
		false,  // immediate
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)


	err = ch.Qos(
		1, 		// prefetch count
		0, 		// prefetch size
		false, 	// global
	)
	failOnError(err, "Failed to set Qos")
}

func bodyForm(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s

}