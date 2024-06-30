package internal

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type message struct {
	Sender   string `json:"sender"`
	Username string `json:"username"`
	Msg      string `json:"msg"`
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
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		message := message{
			Msg:      string(msg),
			Username: c.username,
			Sender:   c.clientId,
		}

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
