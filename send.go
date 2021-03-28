package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func sender(url string) {
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"test",   // name
		"fanout", // type
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	input := make(chan string)
	defer close(input)

	go func() {
		var msg string
		for msg != "exit()" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Messge: ")
			msg, _ := reader.ReadString('\n')
			if msg != "exit()\n" {
				if msg == "burst\n" {
					for i := 1; i < 1000000; i++ {
						input <- fmt.Sprintf("message: %d", i)
					}
				} else {
					input <- msg
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for body := range input {
		err = ch.Publish(
			"test", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}
}
