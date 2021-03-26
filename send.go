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
	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
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
			if msg != "exit()" {
				input <- msg
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for body := range input {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}
}
