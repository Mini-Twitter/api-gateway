package main

import (
	config2 "apigateway/pkg/config"
	"apigateway/pkg/logger"
	"apigateway/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func main() {
	log := logger.NewLogger()

	config := config2.Load()

	service, err := service.NewService(config)
	if err != nil {
		log.Error("Failed to create service")
		return
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	path, err := os.Getwd()
	if err != nil {
		log.Error("Failed to get current working directory")
		return
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
