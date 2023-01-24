package main

import (
	"listener-service/internal/NATS"
	"listener-service/internal/websocket"
)

type Config struct {
}

type Service struct {
	MsgHandler *NATS.MsgHandler
	WS         *websocket.WS
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
	ws := websocket.NewWS()
	msgHandler := NATS.NewMsgHandler()

	service := Service{
		WS:         ws,
		MsgHandler: msgHandler,
	}

	service.Run()
}
