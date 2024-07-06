package internal

import (
	"context"
	"go-chat/internal/db"
	"go-chat/internal/models"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/jackc/pgx/v5"
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

	wsHanlder := NewWsHandler()

	e.GET("/", MustAuth(func(ctx echo.Context) error {
		cookie, _ := ctx.Cookie("user_id")
		rows, err := conn.Query(context.Background(), "SELECT id,username,password,full_name,avatar FROM users WHERE id=$1 LIMIT 1", cookie.Value)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}

		users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.User])
		if err != nil || len(users) == 0 {
			return ctx.String(http.StatusInternalServerError, "server error")
		}
		user := users[0]

		rows, err = conn.Query(context.Background(), "SELECT id,sender_id,message,full_name,created_time FROM messages  ORDER BY created_time DESC LIMIT 10 ")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}

		messages, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Message])
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}

		return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{
			"UserId":   user.ID,
			"Avatar":   user.Avartar,
			"Messages": messages,
		})
	}))

	e.GET("/ws/:id", MustAuth(func(ctx echo.Context) error {
		err := wsHanlder.Serve(ctx, conn)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "server error")
		}
		return nil
	}))

	e.GET("/login", func(ctx echo.Context) error {
		_, err := ctx.Cookie("user_id")
		if err == nil {
			return ctx.Redirect(http.StatusSeeOther, "/")
		}
		return ctx.Render(200, "login.html", nil)
	})

	e.POST("/login", func(ctx echo.Context) error {
		var login models.Login

		err := ctx.Bind(&login)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "bad request")
		}

		rows, err := conn.Query(context.Background(), "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1 and password=$2 LIMIT 1", login.Username, login.Password)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.User])
		if err != nil || len(users) == 0 {
			return ctx.String(http.StatusUnauthorized, "username or password invalid"+err.Error())
		}
		user := users[0]

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
	})

	e.GET("/logout", func(ctx echo.Context) error {
		cookies := ctx.Cookies()

		for _, c := range cookies {
			c.MaxAge = -1
			ctx.SetCookie(c)
		}
		return ctx.Redirect(http.StatusSeeOther, "/login")
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}
