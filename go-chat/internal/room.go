package internal

import (
	"log"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan message
	join    chan *client
	leave   chan *client
	clients map[string]*client
	roomId  string
}

func newRoom(roomId string) *room {
	return &room{
		forward: make(chan message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[string]*client),
		roomId:  roomId,
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			if r.clients[client.clientId] == nil {
				r.clients[client.clientId] = client
			}
			r.sendJoinedOrLeft(client, "joined")
			r.sendCurrUserForNewUser(client)
			log.Println("new client", client.clientId, r.roomId)
		case client := <-r.leave:
			delete(r.clients, client.clientId)
			r.sendJoinedOrLeft(client, "left")
			log.Println("client left")
		case msg := <-r.forward:
			log.Println("new message from ", msg.Username)
			for client := range r.clients {
				if msg.Sender != r.clients[client].clientId {
					r.clients[client].send <- msg
				}
			}
		}
	}
}

func (r *room) sendJoinedOrLeft(client *client, event string) {
	msg := message{
		Sender:   client.clientId,
		Username: client.fullName,
		Msg:      client.fullName + " " + event,
		Type:     event,
	}
	for clientId := range r.clients {
		if clientId != client.clientId {
			r.clients[clientId].send <- msg
		}
	}
}

func (r *room) sendCurrUserForNewUser(newClient *client) {
	for client := range r.clients {
		if client != newClient.clientId {
			msg := message{
				Sender:   r.clients[client].clientId,
				Username: r.clients[client].fullName,
				Type:     "user-list",
			}
			newClient.send <- msg
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}
