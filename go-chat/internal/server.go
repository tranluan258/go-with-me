package internal

import (
	"io"
	"text/template"

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
		return ctx.Render(200, "index.html", nil)
	})
	e.GET("/ws", func(ctx echo.Context) error {
		room.Serve(ctx.Response().Writer, ctx.Request())
		return nil
	})
	go room.run()

	e.Logger.Fatal(e.Start(":8080"))
}
