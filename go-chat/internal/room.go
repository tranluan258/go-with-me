package internal

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan message
	join    chan *client
	leave   chan *client
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
			msg := message{
				Sender:   client.clientId,
				Username: client.username,
				Msg:      client.username + " joined",
				Type:     "notification",
			}
			for client := range r.clients {
				client.send <- msg
			}
			r.clients[client] = true
			log.Println("new client", client.clientId)
		case client := <-r.leave:
			delete(r.clients, client)
			msg := message{
				Sender:   client.clientId,
				Username: client.username,
				Msg:      client.username + " left",
				Type:     "notification",
			}
			for client := range r.clients {
				client.send <- msg
			}
			log.Println("client left")
		case msg := <-r.forward:
			log.Println("new message from ", msg.Username)
			for client := range r.clients {
				if msg.Sender != client.clientId {
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
	cookie, _ := req.Cookie("username")
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		return
	}

	clientId := randomId(10)

	client := &client{
		clientId: clientId,
		username: cookie.Value,
		socket:   socket,
		send:     make(chan message),
		room:     r,
	}
	r.join <- client

	defer func() { r.leave <- client }()

	go client.write()
	client.read()
}

func randomId(length int) string {
	digits := "123456789abcefghjklmnbvcxz"
	res := ""

	for range length {
		res += string(digits[rand.Intn(len(digits))])
	}
	return res
}
