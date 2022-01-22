package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	// Listen for created accounts
	js.Subscribe(userCreated, func(msg *nats.Msg) {
		msg.Ack()

		// Send a welcoming message
		sendEmail("Welcome!", string(msg.Data))

		// TODO: Why MONITOR?
	}, nats.Durable("MONITOR"), nats.ManualAck())

	r := gin.Default()

	log.Fatal(r.Run(":8001"))
}

func sendEmail(subject, address string) {
	from := mail.NewEmail("Boss", "rojasleon.dev@gmail.com")
	to := mail.NewEmail("User", address)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
