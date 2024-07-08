package handlers

import (
	"go-chat/internal/models"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type AuthHander struct {
	db *sqlx.DB
}

func NewLoginHander(db *sqlx.DB) *AuthHander {
	return &AuthHander{
		db: db,
	}
}

func (lh *AuthHander) PostLogin(ctx echo.Context) error {
	var login models.Login

	err := ctx.Bind(&login)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var user models.User

	err = lh.db.Get(&user, "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1 and password=$2", login.Username, login.Password)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "invalid credentials")
	}

	usernameCookie := new(http.Cookie)
	usernameCookie.Name = "full_name"
	usernameCookie.Value = user.FullName
	usernameCookie.Expires = time.Now().Add(24 * time.Hour)
	usernameCookie.HttpOnly = true
	ctx.SetCookie(usernameCookie)

	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = user.ID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)
	return ctx.String(200, "ok")
}

func (lh *AuthHander) LoginGet(ctx echo.Context) error {
	_, err := ctx.Cookie("user_id")
	if err == nil {
		return ctx.Redirect(http.StatusSeeOther, "/")
	}
	return ctx.Render(200, "login.html", nil)
}

func (lh *AuthHander) Logout(ctx echo.Context) error {
	cookies := ctx.Cookies()

	for _, c := range cookies {
		c.MaxAge = -1
		ctx.SetCookie(c)
	}
	return ctx.Redirect(http.StatusSeeOther, "/login")
}
