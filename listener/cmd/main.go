package main

import (
	"listener-service/internal/nats"
	"listener-service/internal/server"
)

type Service struct {
	MsgHandler *nats.MsgHandler
	WS         *server.WS
}

// Entry point of the service
func (s *Service) Run() error {
	go s.WS.ServeHTTP()

	for {
		select {
		case msg := <-s.WS.WSCh:
			s.MsgHandler.Handle(msg)
		}
	}
}

func main() {
	ws := server.NewWS()
	msgHandler := nats.NewMsgHandler()

	service := Service{
		WS:         ws,
		MsgHandler: msgHandler,
	}

	service.Run()
}
