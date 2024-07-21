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
	// TODO: update query get room detail
	return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
		"Messages": messages,
		"UserId":   cookie.Value,
		"Room": map[string]interface{}{
			"ID":   roomId,
			"Name": "No Name",
		},
	})
}
