package ws

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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

func (ws *WsHandler) Serve(c echo.Context, conn *sqlx.DB) error {
	fullName := c.Get("full_name")
	userId := c.Get("user_id")
	socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	roomId := c.Param("id")
	if roomId == "" {
		return err
	}

	existRoom := models.Room{}
	err = conn.Get(&existRoom, "SELECT id FROM rooms WHERE id=$1", roomId)
	if err != nil {
		return c.String(http.StatusNotFound, "not found room")
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
		clientId: userId.(string),
		fullName: fullName.(string),
		socket:   socket,
		send:     make(chan []byte),
		room:     room,
		conn:     conn,
	}
	room.join <- client

	defer func() { room.leave <- client }()

	go client.write()
	client.read()

	return nil
}
