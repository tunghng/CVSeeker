package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Ensure to adjust this based on your CORS policy
	},
}

var (
	connections []*WebSocketConnection
	connMutex   sync.Mutex
)

type WebSocketConnection struct {
	Conn   *websocket.Conn
	send   chan []byte
	closed bool
	mu     sync.Mutex
}

func NewWebSocketConnection(conn *websocket.Conn) *WebSocketConnection {
	wc := &WebSocketConnection{
		Conn:   conn,
		send:   make(chan []byte, 256), // Buffered channel for outgoing messages
		closed: false,
	}
	connMutex.Lock()
	connections = append(connections, wc)
	connMutex.Unlock()
	return wc
}

func (wc *WebSocketConnection) cleanup() {
	wc.mu.Lock()
	if !wc.closed {
		close(wc.send)
		wc.Conn.Close()
		wc.closed = true
	}
	wc.mu.Unlock()
}

func (wc *WebSocketConnection) sendSafe(message []byte) error {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	if wc.closed {
		return fmt.Errorf("attempt to send on closed connection")
	}
	select {
	case wc.send <- message:
		return nil
	default:
		return fmt.Errorf("send buffer is full or closed")
	}
}

func (wc *WebSocketConnection) readPump() {
	defer wc.cleanup()
	wc.Conn.SetReadLimit(512)
	wc.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	wc.Conn.SetPongHandler(func(string) error {
		wc.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := wc.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("WebSocket error: %v\n", err)
			break
		}
		fmt.Printf("Received: %s\n", message)
	}
}

func (wc *WebSocketConnection) writePump() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		wc.cleanup()
	}()
	for {
		select {
		case message, ok := <-wc.send:
			wc.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
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

func HandleWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocketConnection, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	wc := NewWebSocketConnection(conn)
	go wc.writePump()
	go wc.readPump()
	return wc, nil
}

func BroadcastMessage(message []byte) {
	connMutex.Lock()
	defer connMutex.Unlock()
	for i := len(connections) - 1; i >= 0; i-- {
		conn := connections[i]
		if err := conn.sendSafe(message); err != nil {
			fmt.Println("Error sending message:", err)
			connections = append(connections[:i], connections[i+1:]...)
		}
	}
}
