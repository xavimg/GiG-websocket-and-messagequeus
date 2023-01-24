package main

import (
	"log"

	"notifier-service/internal/NATS"
	"notifier-service/internal/websocket"

	"github.com/nats-io/nats.go"
)

type Config struct {
}

type Service struct {
	WSPublisher  *websocket.WSPublisher
	MQSubscriber *NATS.MQSubscriber
}

// Entry point of the service
func (s *Service) Run() error {
	go s.WSPublisher.ServeHTTP()

	s.MQSubscriber.Subscribe("Test")
	go s.WSPublisher.HandleWS()

	for {
		select {
		case msg := <-s.MQSubscriber.MQSCh:
			log.Println("received message: ", msg)
			s.WSPublisher.WSPubCh <- msg.Data
		}
	}
}

func main() {
	ws := websocket.NewWSPublisher()
	mq := NATS.NewMQSubsc(nats.DefaultURL)

	service := Service{
		MQSubscriber: mq,
		WSPublisher:  ws,
	}

	service.Run()
}
