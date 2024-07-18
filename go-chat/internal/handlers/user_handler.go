package handlers

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
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
	cookie, _ := ctx.Cookie("user_id")
	var searhQuery searhQuery

	err := ctx.Bind(&searhQuery)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var users []models.User

	err = uh.db.Select(&users, "SELECT id, full_name,avatar FROM users WHERE id!=$1 AND full_name LIKE $2", cookie.Value, "%"+searhQuery.Search+"%")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	return ctx.Render(http.StatusOK, "user-list", map[string]interface{}{
		"Users": users,
	})
}
