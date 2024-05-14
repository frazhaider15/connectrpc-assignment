package main

import (
	"fmt"

	"github.com/auth/config"
	connectrpc "github.com/auth/connectRPC"
	"github.com/auth/db"
	"github.com/auth/rabbitmq"
)

func main() {
	config.LoadEnvVariables()

	db.ConnectToDb()
	db.SyncDatabase()

	rabbitmq.ConnectRabbitMq()
	
	go connectrpc.StartServer()

	fmt.Println("Server Started")
	select {}
}
