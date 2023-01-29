package server

import (
	"fmt"
	"log"
	"net/http"

	"gig-websockets-messagequeue/notifier/internal/config"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSNotifier struct {
	Message   chan []byte
	WSClients []*websocket.Conn
}

func NewWSNotifier() *WSNotifier {
	msg := make(chan []byte, 64)
	return &WSNotifier{
		Message: msg,
	}
}

func (wsn *WSNotifier) ServeWS() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade connection err: %w", err)
		}

		log.Printf("New client connected to notifier --> %s", websocket.RemoteAddr())
		wsn.WSClients = append(wsn.WSClients, websocket)

		go wsn.handleMessages()
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.Settings.Notifier.Port), nil)
}

func (wsn *WSNotifier) handleMessages() {
	for msgPayload := range wsn.Message {
		for i, client := range wsn.WSClients {
			if err := client.WriteMessage(websocket.TextMessage, msgPayload); err != nil {
				log.Println(err)
				wsn.removeWSClient(i)
				err = client.Close()
				log.Println(err)
			}
		}
	}
}

func (wsp *WSNotifier) removeWSClient(index int) {
	_ = append(wsp.WSClients[:index], wsp.WSClients[index+1:]...)
}
