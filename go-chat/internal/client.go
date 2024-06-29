package internal

import "github.com/gorilla/websocket"

type message struct {
	msg    []byte
	sender int
}

type client struct {
	socket   *websocket.Conn
	send     chan message
	room     *room
	clientId int
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		message := message{
			msg:    msg,
			sender: c.clientId,
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
