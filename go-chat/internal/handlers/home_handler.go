package handlers

import (
	"context"
	"go-chat/internal/models"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

type HomeHandler struct {
	db *pgx.Conn
}

func NewHomeHandler(db *pgx.Conn) *HomeHandler {
	return &HomeHandler{
		db: db,
	}
}

func (hh *HomeHandler) GetHomeTemplate(ctx echo.Context) error {
	cookie, _ := ctx.Cookie("user_id")
	rows, err := hh.db.Query(context.Background(), "SELECT id,username,password,full_name,avatar FROM users WHERE id=$1 LIMIT 1", cookie.Value)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.User])
	if err != nil || len(users) == 0 {
		return ctx.String(http.StatusInternalServerError, "server error")
	}
	user := users[0]

	rows, err = hh.db.Query(context.Background(), "SELECT id,sender_id,message,full_name,created_time FROM messages  ORDER BY created_time DESC LIMIT 10 ")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	messages, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Message])
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{
		"UserId":   user.ID,
		"Avatar":   user.Avartar,
		"Messages": messages,
	})
}
