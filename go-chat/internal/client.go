package internal

import "github.com/gorilla/websocket"

type message struct {
	sender   string
	username string
	msg      []byte
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
			msg:      msg,
			username: c.username,
			sender:   c.clientId,
		}

		c.room.forward <- message
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for message := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, message.msg)
		if err != nil {
			return
		}
	}
}
