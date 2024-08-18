package handlers

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
	var room models.Room
	var messages []models.Message
	roomId := ctx.QueryParam("room_id")
	userId := ctx.Get("user_id")

	if roomId == "" {
		return ctx.String(http.StatusBadRequest, "roomId should not be empty")
	}

	err := mh.db.Select(&messages, "SELECT id,message,sender_id,full_name,created_time FROM messages WHERE room_id=$1 ORDER BY created_time ASC", roomId)
	if err != nil {
		return err
	}

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
      `, userId, roomId)
	ctx.Response().Header().Set("Vary", "HX-Request")
	return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
		"Messages": messages,
		"UserId":   userId,
		"Room":     room,
	})
}
