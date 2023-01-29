package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestHTTPUpgradedToWebSocket(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatal("websocket connection upgrade: %w", err)
		}
	})

	go router.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("Wrong status code: got %v, want %v", status, 200)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func TestWebSocketConnection(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(echo))
	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	if err := ws.WriteMessage(websocket.TextMessage, []byte("test-message")); err != nil {
		t.Fatalf("%v", err)
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if string(p) != "test-message" {
		t.Fatalf("bad message")
	}
}
