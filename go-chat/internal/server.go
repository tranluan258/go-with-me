package internal

import (
	"go-chat/internal/models"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Static("/", "views")
	e.Renderer = t

	room := newRoom()
	e.GET("/", func(ctx echo.Context) error {
		cookie, _ := ctx.Cookie("username")

		if cookie == nil {
			return ctx.Redirect(http.StatusMovedPermanently, "/login")
		}

		return ctx.Render(200, "index.html", nil)
	})
	e.GET("/ws", func(ctx echo.Context) error {
		cookie, _ := ctx.Cookie("username")
		if cookie == nil {
			return nil
		}
		room.Serve(ctx.Response().Writer, ctx.Request())
		return nil
	})
	e.GET("/login", func(ctx echo.Context) error {
		cookie, _ := ctx.Cookie("username")
		if cookie != nil {
			return ctx.Redirect(http.StatusMovedPermanently, "/")
		}
		return ctx.Render(200, "login.html", nil)
	})
	e.POST("/login", func(ctx echo.Context) error {
		var user models.User

		err := ctx.Bind(&user)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "bad request")
		}
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = user.Username
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true
		ctx.SetCookie(cookie)
		return ctx.String(http.StatusOK, "ok")
	})
	go room.run()

	e.Logger.Fatal(e.Start(":8080"))
}
