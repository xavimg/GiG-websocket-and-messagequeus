package nats

import (
	"gig-websockets-messagequeue/listener/internal/config"
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	Nats  *nats.Conn
	Topic string
}

func NewPublisher(url string) *Publisher {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to NATS ")

	return &Publisher{
		Nats:  nc,
		Topic: config.Settings.NATS.Topic,
	}
}

func (m *Publisher) Publish(msg []byte) {
	if err := m.Nats.Publish(m.Topic, msg); err != nil {
		log.Println(err)
	}

	log.Printf("publishing message to NATS: %s", string(msg))
}
