package handlers

import (
	"go-chat/internal/models"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type searhQuery struct {
	Search string `query:"search"`
}

type UserHandler struct {
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (uh *UserHandler) SearchUser(ctx echo.Context) error {
	userId := ctx.Get("user_id")

	var searhQuery searhQuery

	err := ctx.Bind(&searhQuery)
	if err != nil {
		slog.Warn("Missing field", "error", err.Error())
		return ctx.Render(http.StatusBadRequest, "errors", map[string]interface{}{
			"Errors": "Missing field",
		})
	}

	var users []models.User

	err = uh.db.Select(&users, "SELECT id, full_name,avatar FROM users WHERE id!=$1 AND full_name LIKE $2", userId, "%"+searhQuery.Search+"%")
	if err != nil {
		slog.Warn("Error search users", "error", err.Error())
		return ctx.Render(http.StatusBadRequest, "errors", map[string]interface{}{
			"Errors": "Something wrong try again",
		})
	}

	return ctx.Render(http.StatusOK, "user-list", map[string]interface{}{
		"Users": users,
	})
}
