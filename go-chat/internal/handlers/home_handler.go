package handlers

import (
	"go-chat/internal/models"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type HomeHandler struct {
	db *sqlx.DB
}

func NewHomeHandler(db *sqlx.DB) *HomeHandler {
	return &HomeHandler{
		db: db,
	}
}

func (hh *HomeHandler) GetHomeTemplate(ctx echo.Context) error {
	cookie, _ := ctx.Cookie("user_id")
	var user models.User
	var rooms []models.Room

	err := hh.db.Get(&user, "SELECT id,username,full_name,avatar FROM users WHERE id=$1", cookie.Value)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	err = hh.db.Select(&rooms, "SELECT id,name FROM rooms LEFT JOIN user_room ON rooms.id = user_room.room_id WHERE user_room.user_id = $1", cookie.Value)
	if err != nil {
		log.Println(err.Error())
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	return ctx.Render(http.StatusOK, "home.html", map[string]interface{}{
		"UserId":   user.ID,
		"FullName": user.FullName,
		"Avatar":   user.Avatar,
		"Rooms":    rooms,
	})
}
