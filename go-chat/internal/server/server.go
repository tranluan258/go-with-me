package server

import (
	"fmt"
	"go-chat/internal/config"
	"go-chat/internal/db"
	"go-chat/internal/handlers"
	"go-chat/internal/ws"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"cloud.google.com/go/storage"
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
	bucket := config.FirebaseConfig()
	t := &Template{
		templates: template.Must(template.New("base").Funcs(template.FuncMap{"timeAgo": timeAgo}).ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Static("/", "views")
	e.Renderer = t

	// NOTE init route
	intWsRoute(e, db)
	initHomeRoute(e, db)
	initAuthRoute(e, db, bucket)
	initRoomRoute(e, db)
	initUserRoute(e, db)
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
	// TODO: this api this return detail room
	// e.GET("/rooms/:room_id", MustAuth(roomHandler.GetRoomById))
}

func initMessageRoute(e *echo.Echo, db *sqlx.DB) {
	messageHandler := handlers.NewMessageHandler(db)

	e.GET("/messages", MustAuth(messageHandler.GetMessagesByRoom))
}

func initUserRoute(e *echo.Echo, db *sqlx.DB) {
	userHandler := handlers.NewUserHandler(db)

	e.GET("/users", MustAuth(userHandler.SearchUser))
}

func timeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration.Hours() >= 24:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	case duration.Hours() >= 1:
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	case duration.Minutes() >= 1:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	default:
		seconds := int(duration.Seconds())
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}
