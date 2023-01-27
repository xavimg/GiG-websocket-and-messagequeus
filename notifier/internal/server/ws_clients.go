package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSPublisher struct {
	WSPubCh   chan []byte
	WSClients []*websocket.Conn
}

func NewWSPublisher() *WSPublisher {
	wsch := make(chan []byte, 64)
	return &WSPublisher{
		WSPubCh: wsch,
	}
}

func (wsp *WSPublisher) ServeHTTP() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("websocket connection upgrade: %w", err)
		}
		log.Printf("ws client %v connected!\n", websocket.LocalAddr().String())
		wsp.WSClients = append(wsp.WSClients, websocket)
	})

	log.Println("http listen on 3011")
	http.ListenAndServe(":3011", nil)
}

func (wsp *WSPublisher) HandleWS() {
	// send payload to the channel
	for msgPayload := range wsp.WSPubCh {

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

func (wsp *WSPublisher) removeWSClient(index int) {
	_ = append(wsp.WSClients[:index], wsp.WSClients[index+1:]...)
}