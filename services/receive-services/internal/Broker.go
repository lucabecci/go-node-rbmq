package internal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//Broker is a struct for rbmq
type Broker struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

//Message is a struct for the message of the broker
type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

//Initialize is for start and configure the broker
func Initialize(user string, password string) (*Broker, error) {
	connection, err := amqp.Dial("amqp://" + user + ":" + password + "@localhost:5672/")
	if err != nil {
		log.Fatal(err.Error())
	}
	chn, err := connection.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("amqp connected")
	channel := Broker{channel: chn, conn: connection}

	return &channel, nil
}

//SaveTask is a method for save the task of rbmq
func (broker *Broker) SaveTask(queue string) {
	messages, _ := broker.channel.Consume("network", "", true, false, false, false, nil)
	go func() {
		for message := range messages {
			var msg string = string(message.Body)
			data := &Message{}

			result, err := TransformData(msg, data)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(result)
		}
	}()
}

//Close is a func for close the amqp
func (broker *Broker) Close() {
	broker.conn.Close()
	broker.channel.Close()
	fmt.Println("connection and channel closing...")
}

//TransformData is a function for transform the message of the broker
func TransformData(msg string, data *Message) (string, error) {
	err := json.Unmarshal([]byte(msg), data)
	if err != nil {
		return "", err
	}
	json, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
