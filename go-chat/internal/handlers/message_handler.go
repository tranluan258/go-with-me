package handlers

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type MessageHandler struct {
	db *sqlx.DB
}

func NewMessageHandler(db *sqlx.DB) *MessageHandler {
	return &MessageHandler{
		db: db,
	}
}

func (mh *MessageHandler) GetMessagesByRoom(ctx echo.Context) error {
	roomId := ctx.QueryParam("room_id")
	if roomId == "" {
		return ctx.String(http.StatusBadRequest, "roomId should not be empty")
	}

	cookie, _ := ctx.Cookie("user_id")

	var messages []models.Message

	err := mh.db.Select(&messages, "SELECT id,message,sender_id,full_name FROM messages WHERE room_id=$1 ORDER BY created_time DESC", roomId)
	if err != nil {
		return err
	}

	var room models.Room
	mh.db.Get(&room, `
      SELECT 
          r.id AS id,
          CASE
              WHEN r.room_type = 'dm' THEN (SELECT u.full_name FROM users u JOIN user_room ru ON u.id = ru.user_id WHERE ru.room_id = r.id AND u.id != $1)
              ELSE r.name
              END AS name,
          CASE
              WHEN r.room_type = 'dm' THEN (SELECT u.avatar FROM users u JOIN user_room ru ON u.id = ru.user_id WHERE ru.room_id = r.id AND u.id != $1)
              ELSE NULL
              END AS avatar,
          r.room_type
      FROM 
          rooms r
      JOIN 
          user_room ru1 ON r.id = ru1.room_id
      WHERE 
          ru1.user_id = $1
      AND 
          r.room_type IN ('dm', 'group')
      AND 
          r.id=$2;
      `, cookie.Value, roomId)
	return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
		"Messages": messages,
		"UserId":   cookie.Value,
		"Room":     room,
	})
}
