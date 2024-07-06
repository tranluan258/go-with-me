package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
)

type message struct {
	SenderId string `json:"sender_id"`
	FullName string `json:"full_name"`
	Msg      string `json:"msg"`
	Type     string `json:"type"`
}

type client struct {
	socket   *websocket.Conn
	send     chan message
	room     *room
	clientId string
	fullName string
	conn     *pgx.Conn
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
		message.SenderId = c.clientId
		message.FullName = c.fullName
		c.room.forward <- message

		c.insertMessgeToDb(message)
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

func (c *client) insertMessgeToDb(msg message) {
	_, err := c.conn.Exec(context.Background(), "INSERT INTO messages(sender_id,full_name,message) VALUES($1,$2,$3)", msg.SenderId, msg.FullName, msg.Msg)
	if err != nil {
		log.Println("error insert message", err.Error())
		return
	}

	log.Println("insert new message")
}
