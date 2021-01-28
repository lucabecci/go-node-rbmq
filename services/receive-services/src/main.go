package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://golang:go123@localhost:5672/")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer ch.Close()

	chMessages, err := ch.Consume("network", "", true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	noSTP := make(chan bool)

	go func() {
		for messages := range chMessages {
			fmt.Println("message:" + string(messages.Body))
		}
	}()

	<-noSTP
}
