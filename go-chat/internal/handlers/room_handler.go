package handlers

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type RoomHandler struct {
	db *sqlx.DB
}

func NewRoomHandler(db *sqlx.DB) *RoomHandler {
	return &RoomHandler{
		db: db,
	}
}

func (rh *RoomHandler) CreateRoom(ctx echo.Context) error {
	var createRoom models.CreateRoom
	var userInRoom []models.User

	err := ctx.Bind(&createRoom)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid body")
	}

	tx := rh.db.MustBegin()

	var roomId string
	err = tx.Get(&roomId, "INSERT INTO rooms (name) VALUES($1) RETURNING id", createRoom.RoomName)
	if err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	err = tx.Select(&userInRoom, "SELECT id FROM users WHERE id IN $1", createRoom.UserIds)
	if err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	for _, v := range userInRoom {
		tx.MustExec("INSERT INTO user_room (user_id, room_id) VALUES($1, $2)", v.ID, roomId)
	}
	tx.Commit()

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"room_id": roomId,
	})
}

func (rh *RoomHandler) GetRoomById(ctx echo.Context) error {
	roomId := ctx.Param("room_id")
	cookie, _ := ctx.Cookie("user_id")

	var messages []models.Message

	err := rh.db.Select(&messages, "SELECT id,message,sender_id,full_name FROM messages WHERE room_id=$1 ORDER BY created_time DESC", roomId)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
		"Messages": messages,
		"UserId":   cookie.Value,
	})
}
