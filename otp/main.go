package main

import (
	"fmt"

	"github.com/otp/config"
	"github.com/otp/services"

	"github.com/otp/rabbitmq"
)

func main() {
	config.LoadEnvVariables()

	services.ConnectTwilio()
	rabbitmq.ConnectRabbitMq()
	go rabbitmq.ConsumeMessages()

	fmt.Println("Server Started")
	select {}
}
