package main

import (
	"log"

	"notifier-service/internal/broker"
	"notifier-service/internal/server"
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
			log.Println("message through the message queue: ", string(msg.Data))
			s.WSPublisher.WSPubCh <- msg.Data
		}
	}
}

func main() {
	ws := server.NewWSPublisher()
	mq := broker.NewMQSubsc("nats://nats:4222")

	service := Service{
		MQSubscriber: mq,
		WSPublisher:  ws,
	}

	service.Run()
}
