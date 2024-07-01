package internal

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type message struct {
	Sender   string `json:"sender"`
	Username string `json:"username"`
	Msg      string `json:"msg"`
	Type     string `json:"type"`
}

type client struct {
	socket   *websocket.Conn
	send     chan message
	room     *room
	clientId string
	username string
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var message message
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		json.Unmarshal(msg, &message)
		message.Sender = c.clientId
		message.Username = c.username

		c.room.forward <- message
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for message := range c.send {
		data, _ := json.Marshal(message)
		err := c.socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return
		}
	}
}
