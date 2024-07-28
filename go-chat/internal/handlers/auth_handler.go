package handlers

import (
	"context"
	"go-chat/internal/models"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/markbates/goth/gothic"
)

const (
	GOOGLE_PROVDIER   = "google"
	FACEBOOK_PROVIDER = "facebook"
)

type AuthHander struct {
	db     *sqlx.DB
	bucket *storage.BucketHandle
}

func NewLoginHander(db *sqlx.DB, bucket *storage.BucketHandle) *AuthHander {
	return &AuthHander{
		db:     db,
		bucket: bucket,
	}
}

func (ah *AuthHander) LoginPost(ctx echo.Context) error {
	var login models.Login

	err := ctx.Bind(&login)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var user models.User

	err = ah.db.Get(&user, "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1 and password=$2", login.Username, login.Password)
	if err != nil {
		return ctx.Render(http.StatusUnauthorized, "errors", map[string]interface{}{
			"Errors": "Invalid credentials",
		})
	}
	ah.SetCookie(ctx, user.FullName, user.ID)
	ctx.Response().Header().Add("HX-Redirect", "/")
	return ctx.String(http.StatusFound, "login success")
}

func (ah *AuthHander) LoginGet(ctx echo.Context) error {
	_, err := ctx.Cookie("user_id")
	if err == nil {
		return ctx.Redirect(http.StatusSeeOther, "/")
	}
	return ctx.Render(200, "login.html", map[string]interface{}{
		"GoogleUrl":   os.Getenv("GOOGLE_URL"),
		"FacebookUrl": os.Getenv("FACEBOOK_URL"),
	})
}

func (ah *AuthHander) Logout(ctx echo.Context) error {
	cookies := ctx.Cookies()

	for _, c := range cookies {
		c.MaxAge = -1
		ctx.SetCookie(c)
	}
	return ctx.Redirect(http.StatusSeeOther, "/login")
}

func (ah *AuthHander) BeginAuth(ctx echo.Context) error {
	provider := ctx.Param("provider")

	if provider != GOOGLE_PROVDIER && provider != FACEBOOK_PROVIDER {
		return ctx.String(http.StatusBadRequest, "provider not supported")
	}

	q := ctx.Request().URL.Query()
	q.Add("provider", provider)
	ctx.Request().URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(ctx.Response(), ctx.Request())
	return nil
}

func (ah *AuthHander) CompleteAuth(ctx echo.Context) error {
	var existedUser models.User
	oauhtUser, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request())
	if err != nil {
		log.Println(err.Error())
		return ctx.String(http.StatusInternalServerError, "server error")
	}

	isEmpty := ah.db.Get(&existedUser, "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1", oauhtUser.UserID)
	if isEmpty != nil {
		_, err := ah.db.NamedExec("INSERT INTO users (username, password,full_name,avatar) VALUES(:username,:password,:full_name,:avatar)", map[string]interface{}{
			"username":  oauhtUser.UserID,
			"password":  oauhtUser.UserID,
			"full_name": oauhtUser.Email,
			"avatar":    oauhtUser.AvatarURL,
		})
		if err != nil {
			log.Println(err.Error())
			return ctx.String(http.StatusInternalServerError, "server error")
		}

		newUser := models.User{}
		ah.db.Get(&newUser, "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1", oauhtUser.UserID)

		ah.SetCookie(ctx, newUser.FullName, newUser.ID)
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}

	ah.SetCookie(ctx, existedUser.FullName, existedUser.ID)
	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

func (ah *AuthHander) RegisterGet(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "register.html", nil)
}

func (ah *AuthHander) RegisterPost(ctx echo.Context) error {
	var avatar *string = nil
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	fullname := ctx.FormValue("fullname")

	file, err := ctx.FormFile("avatar")
	if err == nil {
		func() {
			src, err := file.Open()
			if err != nil {
				slog.Error("Upload failed", "error", err.Error())
				return
			}
			defer src.Close()

			// Destination
			t := time.Now()

			docName := t.Format("20060102150405") + file.Filename
			wc := ah.bucket.Object(docName).NewWriter(context.Background())
			if _, err = io.Copy(wc, src); err != nil {
				slog.Error("Upload failed", "error", err.Error())
				return
			}
			err = wc.Close()
			if err != nil {
				slog.Error("Upload failed", "error", err.Error())
				return
			}
			url, err := ah.bucket.SignedURL(docName, &storage.SignedURLOptions{
				Method:  "GET",
				Expires: time.Now().AddDate(0, 0, 7),
			})
			if err != nil {
				slog.Error("Get url failed", "error", err.Error())
				return
			}
			avatar = &url
		}()
	}

	var usernameExisted string
	empty := ah.db.Get(&usernameExisted, "SELECT username FROM users WHERE username=$1", username)
	if empty == nil {
		return ctx.Render(http.StatusBadRequest, "errors", map[string]interface{}{
			"Errors": "User already existed",
		})
	}

	_, err = ah.db.Exec("INSERT INTO users(username,password,full_name,avatar) VALUES($1,$2,$3, $4)", username, password, fullname, avatar)
	if err != nil {
		slog.Error("Insert failed", "error", err.Error())
		return ctx.Render(http.StatusBadRequest, "errors", map[string]interface{}{
			"Errors": "Register failed",
		})
	}

	ctx.Response().Header().Add("HX-Redirect", "/login")
	return ctx.String(http.StatusFound, "register success")
}

func (ah *AuthHander) SetCookie(ctx echo.Context, fullName, userId string) {
	usernameCookie := new(http.Cookie)
	usernameCookie.Name = "full_name"
	usernameCookie.Value = fullName
	usernameCookie.Expires = time.Now().Add(24 * time.Hour)
	usernameCookie.HttpOnly = true
	usernameCookie.SameSite = http.SameSiteLaxMode
	usernameCookie.Path = "/"
	ctx.SetCookie(usernameCookie)

	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = userId
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode
	ctx.SetCookie(cookie)
}
