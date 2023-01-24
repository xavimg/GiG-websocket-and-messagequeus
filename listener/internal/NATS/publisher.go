package NATS

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
	log.Printf("publishing to MQ every 3 seconds to simulate data in real-time: %s", string(msg))

	m.NatsClient.Publish(m.Topic, msg)

}

func ConnectMQ() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	return nc
}
