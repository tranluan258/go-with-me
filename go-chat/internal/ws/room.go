package ws

import (
	"fmt"
	"log"
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
			log.Println("new client", client.clientId, r.roomId)
			if r.clients[client.clientId] == nil {
				r.clients[client.clientId] = client
			}
		case client := <-r.leave:
			log.Println("client left")
			delete(r.clients, client.clientId)
		case msg := <-r.forward:
			log.Println("new message from ", msg.FullName)
			for client := range r.clients {
				if msg.SenderId != r.clients[client].clientId {
					message := fmt.Sprintf(`<div id="chat-messages" class="chat-messages" hx-swap-oob="beforeend:#chat-messages"><div class="chat chat-start"> <div class="chat-header"> %s <time class="text-xs opacity-50">Now</time> </div> <div class="chat-bubble chat-bubble-primary">%s</div></div></div>`, msg.FullName, msg.Msg)
					r.clients[client].send <- []byte(message)
				} else {
					message := fmt.Sprintf(`<div id="chat-messages" class="chat-messages" hx-swap-oob="beforeend:#chat-messages"><div class="chat chat-end"> <div class="chat-header"> %s <time class="text-xs opacity-50">Now</time> </div> <div class="chat-bubble chat-bubble-accent">%s</div></div></div>`, "Me", msg.Msg)
					r.clients[client].send <- []byte(message)
				}
			}
		}
	}
}
