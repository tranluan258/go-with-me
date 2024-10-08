package server

import (
	"errors"
	"go-chat/internal/config"
	"go-chat/internal/db"
	"go-chat/internal/handlers"
	"go-chat/internal/helpers"
	"go-chat/internal/ws"
	"io"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"cloud.google.com/go/storage"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func argsfn(kvs ...interface{}) (map[string]interface{}, error) {
	if len(kvs)%2 != 0 {
		return nil, errors.New("args requires even number of arguments.")
	}
	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		s, ok := kvs[i].(string)
		if !ok {
			return nil, errors.New("even args to args must be strings.")
		}
		m[s] = kvs[i+1]
	}
	return m, nil
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		slog.Error("error render html: ", "err", err.Error())
		return err
	}
	return nil
}

func Init() {
	db := db.InitDb()
	bucket := config.FirebaseConfig()
	t := &Template{
		templates: template.Must(template.New("base").Funcs(template.FuncMap{"timeAgo": helpers.TimeAgo, "args": argsfn}).ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Static("/", "views")
	e.Renderer = t

	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	gothic.Store = store

	e.Use(session.Middleware(store))

	intWsRoute(e, db)
	initHomeRoute(e, db)
	initAuthRoute(e, db, bucket)
	initRoomRoute(e, db)
	initUserRoute(e, db)
	initMessageRoute(e, db)

	e.Logger.Fatal(e.Start(":" + "8080"))
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

func initAuthRoute(e *echo.Echo, db *sqlx.DB, bucket *storage.BucketHandle) {
	loginHandler := handlers.NewLoginHander(db, bucket)

	e.GET("/auth/:provider", loginHandler.BeginAuth)
	e.GET("/auth/:provider/callback", loginHandler.CompleteAuth)
	e.GET("/register", loginHandler.RegisterGet)
	e.POST("/register", loginHandler.RegisterPost)
	e.GET("/login", loginHandler.LoginGet)
	e.POST("/login", loginHandler.LoginPost)
	e.GET("/logout", loginHandler.Logout)
}

func initRoomRoute(e *echo.Echo, db *sqlx.DB) {
	roomHandler := handlers.NewRoomHandler(db)

	e.POST("/rooms", MustAuth(roomHandler.CreateRoom))
	e.GET("/rooms/dm-room", MustAuth(roomHandler.GetDMRoom))
}

func initMessageRoute(e *echo.Echo, db *sqlx.DB) {
	messageHandler := handlers.NewMessageHandler(db)

	e.GET("/messages", MustAuth(messageHandler.GetMessagesByRoom))
}

func initUserRoute(e *echo.Echo, db *sqlx.DB) {
	userHandler := handlers.NewUserHandler(db)

	e.GET("/users", MustAuth(userHandler.SearchUser))
}
