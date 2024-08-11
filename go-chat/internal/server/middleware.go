package server

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sess, err := session.Get("session_id", ctx)
		if err != nil {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}

		userID, ok := sess.Values["user_id"]
		if !ok {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		if v, ok := userID.(string); ok {
			ctx.Set("user_id", v)
		}

		fullName, ok := sess.Values["full_name"]
		if !ok {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		if v, ok := fullName.(string); ok {
			ctx.Set("full_name", v)
		}

		return next(ctx)
	}
}
