package internal

import (
	"github.com/labstack/echo"
)

type WsHandler struct {
	rooms map[string]*room
}

func NewWsHandler() *WsHandler {
	return &WsHandler{
		rooms: make(map[string]*room),
	}
}

func (ws *WsHandler) Serve(c echo.Context) {
	// connect to ws
	cookie, _ := c.Cookie("username")
	socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return
	}

	roomId := c.Param("id")

	if roomId == "" {
		return
	}

	var room *room

	if ws.rooms[roomId] == nil {
		room = newRoom(roomId)
		ws.rooms[roomId] = room
		go room.run()
	} else {
		room = ws.rooms[roomId]
	}

	clientId := randomId(10)

	client := &client{
		clientId: clientId,
		username: cookie.Value,
		socket:   socket,
		send:     make(chan message),
		room:     room,
	}
	room.join <- client

	defer func() { room.leave <- client }()

	go client.write()
	client.read()
}
