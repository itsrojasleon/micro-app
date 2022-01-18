package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

const (
	streamName     = "USERS"
	streamSubjects = "USERS.*"
)

var NatsConn *nats.Conn

func ConnectToNATS() {
	NatsConn, _ = nats.Connect(os.Getenv("NATS_URL"))

	js, err := GetJetstream(NatsConn)
	if err != nil {
		log.Fatal(err)
	}

	err = createStream(js)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to NATS")
}

func GetJetstream(nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))

	if err != nil {
		return nil, err
	}

	return js, nil
}

func createStream(js nats.JetStreamContext) error {
	stream, _ := js.StreamInfo(streamName)

	if stream == nil {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
