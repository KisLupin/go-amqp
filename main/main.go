package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("AMPQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Success")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(q)
	i := 0
	for i < 10 {
		i++
		err = ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("hello world"),
			},
		)

		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}

	fmt.Println("success publish")
}
