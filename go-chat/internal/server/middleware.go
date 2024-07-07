package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Cookie("user_id")
		if err != nil {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		return next(ctx)
	}
}
