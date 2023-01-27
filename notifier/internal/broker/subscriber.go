package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MQSubscriber struct {
	Nats  *nats.Conn
	MQSCh chan *nats.Msg
}

func NewMQSubsc(natsURL string) *MQSubscriber {
	// Connect to a server
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to NATS Server: %s\n", natsURL)

	mqch := make(chan *nats.Msg)
	return &MQSubscriber{
		Nats:  nc,
		MQSCh: mqch,
	}
}

func (mqs *MQSubscriber) Subscribe(topic string) *nats.Subscription {
	sub, err := mqs.Nats.ChanSubscribe(topic, mqs.MQSCh)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("subscribed to %s\n", topic)
	return sub
}

func (mqs *MQSubscriber) Unsubscribe(sub *nats.Subscription) {
	_ = sub.Unsubscribe()
	mqs.Nats.Close()
}
