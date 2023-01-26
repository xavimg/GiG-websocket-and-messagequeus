package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	NatsClient *nats.Conn
	Topic      string
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		NatsClient: ConnectMQ(),
		Topic:      "Test",
	}
}

func (m *MsgHandler) Handle(msg []byte) {
	log.Printf("publishing message to mq: %s", string(msg))
	m.NatsClient.Publish(m.Topic, msg)
}

func ConnectMQ() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	return nc
}
