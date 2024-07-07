package ws

import (
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

type WsHandler struct {
	rooms map[string]*room
}

func NewWsHandler() *WsHandler {
	return &WsHandler{
		rooms: make(map[string]*room),
	}
}

func (ws *WsHandler) Serve(c echo.Context, conn *pgx.Conn) error {
	fullName, _ := c.Cookie("full_name")
	clientId, _ := c.Cookie("user_id")
	socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	roomId := c.Param("id")
	if roomId == "" {
		return err
	}

	var room *room

	if ws.rooms[roomId] == nil {
		room = newRoom(roomId)
		ws.rooms[roomId] = room
		go room.run() // NOTE: run room one times
	} else {
		room = ws.rooms[roomId]
	}

	client := &client{
		clientId: clientId.Value,
		fullName: fullName.Value,
		socket:   socket,
		send:     make(chan message),
		room:     room,
		conn:     conn,
	}
	room.join <- client

	defer func() { room.leave <- client }()

	go client.write()
	client.read()

	return nil
}