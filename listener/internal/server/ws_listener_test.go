package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
