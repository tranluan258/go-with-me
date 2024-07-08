package server

import (
	"go-chat/internal/db"
	"go-chat/internal/handlers"
	"go-chat/internal/ws"
	"io"
	"net/http"
	"text/template"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {
	conn := db.InitDb()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Static("/", "views")
	e.Renderer = t

	// NOTE init route
	intWsRoute(e, conn)
	initHomeRoute(e, conn)
	initLoginRoute(e, conn)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func intWsRoute(e *echo.Echo, conn *sqlx.DB) {
	wsHanlder := ws.NewWsHandler()

	e.GET("/ws/:id", MustAuth(func(ctx echo.Context) error {
		err := wsHanlder.Serve(ctx, conn)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}
		return nil
	}))
}

func initHomeRoute(e *echo.Echo, conn *sqlx.DB) {
	homeHandler := handlers.NewHomeHandler(conn)

	e.GET("/", MustAuth(homeHandler.GetHomeTemplate))
}

func initLoginRoute(e *echo.Echo, conn *sqlx.DB) {
	loginHandler := handlers.NewLoginHander(conn)

	e.GET("/login", loginHandler.LoginGet)
	e.POST("/login", loginHandler.PostLogin)
	e.GET("/logout", loginHandler.Logout)
}
