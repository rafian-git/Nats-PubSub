package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const limit = 100

func main() {
	subject := "test"
	// Connect to nast server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("publisher will publish on subject : %s\n", subject)

	for i := 0; i < limit; i++ {
		log.Printf("Publishing message no: %v\n", i+1)

		if err := nc.Publish(subject, []byte(string(i+1))); err != nil {
			log.Fatal("error publishing:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
