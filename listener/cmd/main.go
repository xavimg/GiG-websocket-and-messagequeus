package main

import (
	"gig-websockets-messagequeue/listener/internal/config"
	"gig-websockets-messagequeue/listener/internal/nats"
	"gig-websockets-messagequeue/listener/internal/server"
	"log"
)

type Service struct {
	WSL *server.WSListener
	Pub *nats.Publisher
}

func (s *Service) Run() error {
	go s.WSL.ServeWS()

	for {
		select {
		case msg := <-s.WSL.Message:
			s.Pub.Publish(msg)
		}
	}
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	wsl := server.NewWSListener()
	pub := nats.NewPublisher(config.Settings.NATS.URL)

	service := Service{
		WSL: wsl,
		Pub: pub,
	}

	service.Run()
}
