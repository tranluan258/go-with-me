package handlers

import (
	"go-chat/internal/models"
	"log/slog"
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

	err := ctx.Bind(&createRoom)
	if err != nil {
		slog.Error(err.Error())
		return ctx.String(http.StatusBadRequest, "invalid body")
	}

	cookie, _ := ctx.Cookie("user_id")
	createRoom.UserIds = append(createRoom.UserIds, cookie.Value)

	tx := rh.db.MustBegin()

	var roomId string
	err = tx.Get(&roomId, "INSERT INTO rooms (name) VALUES($1) RETURNING id", createRoom.RoomName)
	if err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	for _, v := range createRoom.UserIds {
		tx.MustExec("INSERT INTO user_room (user_id, room_id) VALUES($1, $2)", v, roomId)
	}
	tx.Commit()

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"room_id":   roomId,
		"room_name": createRoom.RoomName,
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

func (rh *RoomHandler) GetDMRoom(ctx echo.Context) error {
	user2Id := ctx.QueryParam("user_id")

	if user2Id == "" {
		return ctx.String(http.StatusBadRequest, "user2Id should not be empty")
	}
	cookie, _ := ctx.Cookie("user_id")

	var existedRoom models.Room

	err := rh.db.Get(&existedRoom,
		`SELECT ru1.room_id AS id, r.name as name FROM user_room   ru1
     JOIN user_room ru2 ON ru1.room_id = ru2.room_id
     JOIN rooms r ON ru1.room_id = r.id
     WHERE ru1.user_id = $1
     AND ru2.user_id = $2
     AND r.room_type = 'dm'
     GROUP BY ru1.room_id,r.name`, user2Id, cookie.Value)
	if err != nil {
		slog.Error("Empty rooms", "debug", err.Error())
		return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
			"Messages": []models.Message{},
			"UserId":   cookie.Value,
			"Room": map[string]interface{}{
				"ID":   user2Id,
				"Name": "No Name",
			},
		})
	}

	var messages []models.Message

	err = rh.db.Select(&messages, "SELECT id,message,sender_id,full_name FROM messages WHERE room_id=$1 ORDER BY created_time DESC", existedRoom.ID)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "messages", map[string]interface{}{
		"Messages": messages,
		"UserId":   cookie.Value,
		"Room":     existedRoom,
	})
}
