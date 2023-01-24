package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WS struct {
	WSCh chan []byte
}

func NewWS() *WS {
	wsch := make(chan []byte)
	return &WS{
		WSCh: wsch,
	}
}

func (ws *WS) ServeHTTP() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("http upgrades to websocket connection error: %w", err)
		}
		log.Printf("client from %s connected throught WS ", r.RemoteAddr)
		ws.listenWS(websocket)
	})

	log.Println("Server start on 8080")
	http.ListenAndServe(":8080", nil)
}

func (ws *WS) listenWS(conn *websocket.Conn) {
	for {
		// read a message
		messageType, messageContent, err := conn.ReadMessage()
		timeReceive := time.Now()
		if err != nil {
			log.Println(err)
			return
		}

		// send payload to the channel
		ws.WSCh <- messageContent

		// reponse message
		messageResponse := fmt.Sprintf("Your message is: %s. Time received : %v", messageContent, timeReceive)

		if err := conn.WriteMessage(messageType, []byte(messageResponse)); err != nil {
			log.Println(err)
			return
		}

	}
}
