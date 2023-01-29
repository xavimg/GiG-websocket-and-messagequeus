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

func (wsp *WSNotifier) ServeWS() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade connection err: %w", err)
		}
		log.Printf("New client connected to listener --> %s", websocket.RemoteAddr())
		wsp.WSClients = append(wsp.WSClients, websocket)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.Settings.Notifier.Port), nil)
}

func (wsp *WSNotifier) HandleWS() {
	for msgPayload := range wsp.Message {
		for i, client := range wsp.WSClients {
			if err := client.WriteMessage(websocket.TextMessage, msgPayload); err != nil {
				log.Println(err)
				wsp.removeWSClient(i)
				err = client.Close()
				log.Println(err)
			}
		}
	}
}

func (wsp *WSNotifier) removeWSClient(index int) {
	_ = append(wsp.WSClients[:index], wsp.WSClients[index+1:]...)
}
