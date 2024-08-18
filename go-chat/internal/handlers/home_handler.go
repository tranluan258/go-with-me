package handlers

import (
	"go-chat/internal/models"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
	userId := ctx.Get("user_id")

	var user models.User
	var rooms []models.Room

	err := hh.db.Get(&user, "SELECT id,username,full_name,avatar FROM users WHERE id=$1", userId)
	if err != nil {
		return ctx.String(http.StatusNotFound, "not found user")
	}

	err = hh.db.Select(&rooms, `
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
          r.room_type IN ('dm', 'group');`,
		userId)
	if err != nil {
		slog.Error("Error not found room", "error", err.Error())
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	return ctx.Render(http.StatusOK, "home.html", map[string]interface{}{
		"UserId":   user.ID,
		"FullName": user.FullName,
		"Avatar":   user.Avatar,
		"Rooms":    rooms,
	})
}
