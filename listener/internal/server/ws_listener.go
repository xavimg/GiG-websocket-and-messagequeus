package server

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
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("websocket connection upgrade: %w", err)
		}
		log.Println("ws connected!")
		go ws.listenWS(websocket)
	})

	log.Println("http listen on 3010")
	http.ListenAndServe(":3010", nil)
}

func (ws *WS) listenWS(conn *websocket.Conn) {
	for {
		messageType, messageContent, err := conn.ReadMessage()
		timeReceive := time.Now()
		if err != nil {
			log.Println(err)
			return
		}

		ws.WSCh <- messageContent

		messageResponse := fmt.Sprintf("Your message is: %s. Time received : %v", messageContent, timeReceive)

		if err := conn.WriteMessage(messageType, []byte(messageResponse)); err != nil {
			log.Println(err)
			return
		}

	}
}
