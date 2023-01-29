package server

import (
	"fmt"
	"gig-websockets-messagequeue/listener/internal/config"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSListener struct {
	Message chan []byte
}

func NewWSListener() *WSListener {
	msg := make(chan []byte)
	return &WSListener{
		Message: msg,
	}
}

func (wsl *WSListener) ServeWS() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("upgrade connection err: %w", err)
		}

		log.Println(fmt.Sprintf("New client connected to listener --> %s", websocket.RemoteAddr()))

		go wsl.handleMessages(websocket)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.Settings.Listener.Port), nil)
}

func (wsl *WSListener) handleMessages(conn *websocket.Conn) {
	for {
		messageType, messageContent, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		wsl.Message <- messageContent

		if err := conn.WriteMessage(messageType, []byte(fmt.Sprintf("Your message is: %s. Time received : %v", messageContent, time.Now()))); err != nil {
			log.Println(err)
			return
		}

	}
}
