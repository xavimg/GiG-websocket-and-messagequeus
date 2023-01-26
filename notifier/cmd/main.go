package main

import (
	"log"

	"notifier-service/internal/broker"
	"notifier-service/internal/server"

	"github.com/nats-io/nats.go"
)

type Service struct {
	WSPublisher  *server.WSPublisher
	MQSubscriber *broker.MQSubscriber
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
	ws := server.NewWSPublisher()
	mq := broker.NewMQSubsc(nats.DefaultURL)

	service := Service{
		MQSubscriber: mq,
		WSPublisher:  ws,
	}

	service.Run()
}
