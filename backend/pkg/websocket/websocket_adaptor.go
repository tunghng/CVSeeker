package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type IWebSocketClient interface {
	StartWebSocketClient() error
}

type webSocketClient struct {
	conn *websocket.Conn
}

func NewWebSocketClient() IWebSocketClient {
	return &webSocketClient{}
}

func (w *webSocketClient) StartWebSocketClient() error {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/websocket"}
	var err error
	w.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	go func() {
		defer w.conn.Close()
		for {
			_, message, err := w.conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				break
			}
			log.Printf("Received: %s", message)
		}
	}()
	return nil
}
