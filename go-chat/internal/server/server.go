package server

import (
	"go-chat/internal/db"
	"go-chat/internal/handlers"
	"go-chat/internal/ws"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func Init() {
	db := db.InitDb()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Static("/", "views")
	e.Renderer = t

	// NOTE init route
	intWsRoute(e, db)
	initHomeRoute(e, db)
	initLoginRoute(e, db)
	initRoomRoute(e, db)
	initMessageRoute(e, db)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func intWsRoute(e *echo.Echo, db *sqlx.DB) {
	wsHanlder := ws.NewWsHandler()

	e.GET("/ws/:id", MustAuth(func(ctx echo.Context) error {
		err := wsHanlder.Serve(ctx, db)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}
		return nil
	}))
}

func initHomeRoute(e *echo.Echo, db *sqlx.DB) {
	homeHandler := handlers.NewHomeHandler(db)

	e.GET("/", MustAuth(homeHandler.GetHomeTemplate))
}

func initLoginRoute(e *echo.Echo, db *sqlx.DB) {
	loginHandler := handlers.NewLoginHander(db)

	e.GET("/auth/:provider", loginHandler.BeginAuth)
	e.GET("/auth/:provider/callback", loginHandler.CompleteAuth)
	e.GET("/login", loginHandler.LoginGet)
	e.POST("/login", loginHandler.PostLogin)
	e.GET("/logout", loginHandler.Logout)
}

func initRoomRoute(e *echo.Echo, db *sqlx.DB) {
	roomHandler := handlers.NewRoomHandler(db)

	e.POST("/rooms", MustAuth(roomHandler.CreateRoom))
	// TODO: this api this return detail room
	// e.GET("/rooms/:room_id", MustAuth(roomHandler.GetRoomById))
}

func initMessageRoute(e *echo.Echo, db *sqlx.DB) {
	messageHandler := handlers.NewMessageHandler(db)

	e.GET("/messages", MustAuth(messageHandler.GetMessagesByRoom))
}
