package main

import (
	"fmt"
	"github.com/go-stomp/stomp/v3"
	"log"
)

func main() {
	// Connect to ActiveMQ
	conn, err := stomp.Dial("tcp",
		"localhost:61616",
		stomp.ConnOpt.Login("admin", "admin"),
		stomp.ConnOpt.HeartBeat(0, 0))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Disconnect()

	// Subscribe to a destination (queue or topic)
	sub, err := conn.Subscribe("queue-teste", stomp.AckAuto)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Send a message
	err = conn.Send("queue-teste", "text/plain", []byte("Hello, ActiveMQ!"), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages
	for {
		msg := <-sub.C
		fmt.Printf("Received message: %s\n", msg.Body)

		// Acknowledge the message
		err := conn.Ack(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
