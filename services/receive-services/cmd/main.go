package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lucabecci/go-node-rbmq/services/receive-services/internal"
)

func main() {
	fmt.Println("APP WORKING")
	user, password, queue := loadingENV()
	fmt.Println(user, password)
	//timeout for docker
	time.Sleep(15 * time.Second)

	broker, err := internal.Initialize("guest", "guest")

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
