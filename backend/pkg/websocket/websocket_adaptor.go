package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Adjust the CheckOrigin function to a more secure version based on your requirements.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketConnection wraps a WebSocket connection and a channel for sending messages.
type WebSocketConnection struct {
	Conn *websocket.Conn
	send chan []byte
}

// NewWebSocketConnection creates a new WebSocketConnection instance.
func NewWebSocketConnection(conn *websocket.Conn) *WebSocketConnection {
	return &WebSocketConnection{
		Conn: conn,
		send: make(chan []byte, 256), // Buffered channel for outgoing messages
	}
}

// readPump handles incoming messages from the WebSocket connection.
func (wc *WebSocketConnection) readPump() {
	defer func() {
		wc.Conn.Close()
	}()
	wc.Conn.SetReadLimit(512)
	wc.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	wc.Conn.SetPongHandler(func(string) error { wc.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
	for {
		_, _, err := wc.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Log error or handle it appropriately
			}
			break
		}

	}
}

// writePump handles outgoing messages to the WebSocket connection.
func (wc *WebSocketConnection) writePump() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		wc.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-wc.send:
			wc.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				wc.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := wc.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			wc.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := wc.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// HandleWebSocket upgrades the HTTP server connection to the WebSocket protocol.
func HandleWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocketConnection, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	wc := NewWebSocketConnection(conn)
	// Start reading and writing routines.
	go wc.writePump()
	go wc.readPump()
	return wc, nil
}
