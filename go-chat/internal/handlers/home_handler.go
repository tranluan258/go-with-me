package handlers

import (
	"go-chat/internal/models"
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
	var listFiends []models.User

	err := hh.db.Get(&user, "SELECT id,username,full_name,avatar FROM users WHERE id=$1", cookie.Value)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	err = hh.db.Select(&listFiends, "SELECT id,username,full_name,avatar FROM users WHERE id!=$1", cookie.Value)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	return ctx.Render(http.StatusOK, "home.html", map[string]interface{}{
		"UserId":   user.ID,
		"Avatar":   user.Avatar,
		"Messages": make([]models.Message, 0),
		"Friends":  listFiends,
	})
}
