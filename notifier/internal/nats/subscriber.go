package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	Nats    *nats.Conn
	Message chan *nats.Msg
}

func NewSubscriber(url string) *Subscriber {
	client, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to NATS")

	msg := make(chan *nats.Msg)
	return &Subscriber{
		Nats:    client,
		Message: msg,
	}
}

func (mqs *Subscriber) Subscribe(topic string) *nats.Subscription {
	sub, err := mqs.Nats.ChanSubscribe(topic, mqs.Message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("subscribed to %s\n", topic)
	return sub
}

func (mqs *Subscriber) Unsubscribe(sub *nats.Subscription) {
	_ = sub.Unsubscribe()
	mqs.Nats.Close()
}
