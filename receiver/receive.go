package receiver

import (
	"log"

	"github.com/streadway/amqp"
)

// Check and return value for each amqp call
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Connect to the rabbitMQ server
func connectToRabbitMQ() {
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()
}

func createChannel() {
	ch, err := conn.Channel()
	failOnError(err, )
}