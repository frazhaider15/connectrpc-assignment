package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/otp/services"
)

func ConsumeMessages() {
	// Consume messages from the queue
	msgs, err := Channel.Consume(
		"verification", // Queue name
		"",             // Consumer name
		true,           // Auto-ack (acknowledged automatically)
		false,          // Exclusive
		false,          // No local
		false,          // No wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("Received message:", string(d.Body))
			var otplog services.OtpLog
			err = json.Unmarshal(d.Body, &otplog)
			if err != nil {
				log.Fatal(err)
			}
			services.SendOtp(otplog)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
