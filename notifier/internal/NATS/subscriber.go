package NATS

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MQSubscriber struct {
	MQSCh chan *nats.Msg
	conn  *nats.Conn
}

func NewMQSubsc(url string) *MQSubscriber {
	// Connect to a server
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to NATS Server: %s\n", url)

	mqch := make(chan *nats.Msg)
	return &MQSubscriber{
		MQSCh: mqch,
		conn:  nc,
	}
}

func (mqs *MQSubscriber) Subscribe(topic string) *nats.Subscription {
	sub, err := mqs.conn.ChanSubscribe(topic, mqs.MQSCh)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("subscribed to %s\n", topic)
	return sub
}

func (mqs *MQSubscriber) Unsubscribe(sub *nats.Subscription) {
	_ = sub.Unsubscribe()
	mqs.conn.Close()
}
