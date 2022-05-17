package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	"log"
	"runtime"
)

func main() {
	subject := "test"
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("Subscriberis listening on subject : %s\n", subject)
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Println("Received message no. ", m.Data)
	})
	if err != nil {
		log.Fatalf("error subscribing to %s:%v\n", subject, err)
	}

	// Exit the main goroutine but allow subscribe to continue running
	runtime.Goexit()
}
