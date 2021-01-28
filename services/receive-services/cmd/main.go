package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lucabecci/go-node-rbmq/services/receive-services/internal"
)

func main() {
	user, password, queue := loadingENV()
	broker, err := internal.Initialize(user, password)

	if err != nil {
		log.Fatal(err.Error())
	}

	broker.SaveTask(queue)

	//If the developer use the Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	broker.Close()
}

func loadingENV() (string, string, string) {
	user := os.Getenv("USER_BROKER")
	password := os.Getenv("PASSWORD_BROKER")
	queue := os.Getenv("QUEUE_BROKER")

	return user, password, queue
}
