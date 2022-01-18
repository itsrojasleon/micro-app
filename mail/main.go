package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

const (
	// Subjects
	userCreated = "USERS.created"
)

func main() {
	nc, _ := nats.Connect("nats://nats:4222")
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.Subscribe(userCreated, func(msg *nats.Msg) {
		msg.Ack()

		// TODO: Send emails!
		email := string(msg.Data)

		fmt.Println("Received MSG: ", email)
		// TODO: Why MONITOR?
	}, nats.Durable("MONITOR"), nats.ManualAck())

	r := gin.Default()

	log.Fatal(r.Run(":8001"))
}

// func createConsumer(js nats.JetStreamContext) {
// 	_, err := js.AddConsumer(stream, &nats.ConsumerConfig{
// 		Durable:   consumer,
// 		AckPolicy: nats.AckAllPolicy,
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
