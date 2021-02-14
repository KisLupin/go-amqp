package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("consumer app")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer ch.Close()

	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Receive msg : %s\n", d.Body)
		}
	}()

	fmt.Println("success connect to RabbitMQ")
	fmt.Println("[*] waiting for msg")
	<- forever
}
