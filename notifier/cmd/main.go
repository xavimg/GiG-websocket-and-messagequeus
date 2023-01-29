package main

import (
	"fmt"
	"log"

	"gig-websockets-messagequeue/notifier/internal/config"
	"gig-websockets-messagequeue/notifier/internal/nats"
	"gig-websockets-messagequeue/notifier/internal/server"
)

type Service struct {
	WSN *server.WSNotifier
	Sub *nats.Subscriber
}

func (s *Service) Run() error {
	go s.WSN.ServeWS()

	fmt.Println(config.Settings.NATS.Topic)
	s.Sub.Subscribe(config.Settings.NATS.Topic)
	go s.WSN.HandleWS()

	for {
		select {
		case msg := <-s.Sub.Message:
			log.Println("message from the message queue: ", string(msg.Data))
			s.WSN.Message <- msg.Data
		}
	}
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	wsn := server.NewWSNotifier()
	sub := nats.NewSubscriber(config.Settings.NATS.URL)

	service := Service{
		WSN: wsn,
		Sub: sub,
	}

	service.Run()
}
