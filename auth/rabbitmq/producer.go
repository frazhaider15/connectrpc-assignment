package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func PublishMessage(msg interface{}) error {

	byteMsg ,err := json.Marshal(&msg)
	if err !=nil{
		return err
	}
	
	// Publish message to the queue
	err = Channel.Publish(
		"",             // Exchange (empty string for default exchange)
		"verification", // Routing key (matches the queue name)
		false,          // Mandatory (return error if undeliverable)
		false,          // Immediate (wait for confirmation before returning)
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        byteMsg,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully!")
	return nil
}
