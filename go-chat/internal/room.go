package internal

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var doomId int = 1

type room struct {
	forward chan message

	join chan *client

	leave chan *client

	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			log.Println("new client", client.clientId)
		case client := <-r.leave:
			delete(r.clients, client)
			log.Println("client left")
		case msg := <-r.forward:
			for client := range r.clients {
				if msg.sender != client.clientId {
					log.Println("new message")
					client.send <- msg
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

func (r *room) Serve(w http.ResponseWriter, req *http.Request) {
	// connect to ws
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		return
	}
	doomId = doomId + 1

	client := &client{
		clientId: doomId,
		socket:   socket,
		send:     make(chan message),
		room:     r,
	}
	r.join <- client

	defer func() { r.leave <- client }()

	go client.write()
	client.read()
}
