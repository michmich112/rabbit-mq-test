package main

import (
	"fmt"
	"log"
	"os"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	args := os.Args[1:]

	if args[0] == "send" {
		sender("amqp://guest:guest@localhost:5672/")
	} else if args[0] == "receive" {
		receiver("amqp://guest:guest@localhost:5672/")
	} else {
		fmt.Println("Argument send or receive required")
	}

}
