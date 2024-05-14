package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func ConnectRabbitMq() {
	var err error
	// RabbitMQ connection details
	url := os.Getenv("RABBITMQ_URL")

	// Connect to RabbitMQ
	Conn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	// Open a channel
	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	//defer Channel.Close()

	// Declare a queue (optional, RabbitMQ will create it if it doesn't exist)
	_, err = Channel.QueueDeclare(
		"verification", // Queue name
		true,           // Durable (will survive server restarts)
		false,          // Delete when unused
		false,
		false, // Exclusive (only this connection can access)
		nil,   // Arguments
	)
	if err != nil {
		log.Fatal(err)
	}
}
