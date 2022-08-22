package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu               sync.Mutex
	connectedClients map[*client]struct{}
}

func NewHub() *Hub {
	return &Hub{
		mu:               sync.Mutex{},
		connectedClients: make(map[*client]struct{}),
	}
}

func (h *Hub) AddClient(cl *client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.connectedClients[cl] = struct{}{}
}

func (h *Hub) RemoveClient(cl *client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.connectedClients, cl)
}

func (h *Hub) Send(data interface{}) {
	h.mu.Lock()
	clients := h.connectedClients
	h.mu.Unlock()

	for cl := range clients {
		if cl.IsAlive() {
			cl.out <- data
		}
	}
}

type client struct {
	conn  *websocket.Conn
	alive bool
	hub   *Hub

	out chan interface{}
}

func NewClient(conn *websocket.Conn, hub *Hub) *client {
	return &client{
		conn:  conn,
		alive: true,
		hub:   hub,
		out:   make(chan interface{}, 10),
	}
}

func (c *client) IsAlive() bool {
	return c.alive
}

func (c *client) Watch() {
	defer func() {
		c.alive = false
		c.conn.Close()
		c.hub.RemoveClient(c)
	}()

	for {
		select {
		case m := <-c.out:
			err := c.conn.WriteJSON(m)
			if err != nil {
				return
			}
		}
	}
}
