package websocket

import (
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	clients map[*WebSocketConnection]bool

	// Broadcast incoming messages to all clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *WebSocketConnection

	// Unregister requests from clients.
	unregister chan *WebSocketConnection

	// Mutex to protect access to clients map
	lock sync.Mutex
}

// NewHub creates a new Hub
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *WebSocketConnection),
		unregister: make(chan *WebSocketConnection),
		clients:    make(map[*WebSocketConnection]bool),
	}
}

// Run starts handling the hub's operations including registering, unregistering, and broadcasting.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.lock.Lock()
			h.clients[client] = true
			h.lock.Unlock()
		case client := <-h.unregister:
			h.lock.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send) // Ensure the writePump goroutine stops
			}
			h.lock.Unlock()
		case message := <-h.broadcast:
			h.lock.Lock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.lock.Unlock()
		}
	}
}

// BroadcastMessage sends a message to all connected clients.
func (h *Hub) BroadcastMessage(message []byte) {
	h.broadcast <- message
}
