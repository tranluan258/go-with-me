package handlers

import (
	"go-chat/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/markbates/goth/gothic"
)

const (
	GOOGLE_PROVDIER   = "google"
	FACEBOOK_PROVIDER = "facebook"
)

type AuthHander struct {
	db *sqlx.DB
}

func NewLoginHander(db *sqlx.DB) *AuthHander {
	return &AuthHander{
		db: db,
	}
}

func (ah *AuthHander) PostLogin(ctx echo.Context) error {
	var login models.Login

	err := ctx.Bind(&login)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var user models.User

	err = ah.db.Get(&user, "SELECT id,username,password,full_name,avatar FROM users WHERE username=$1 and password=$2", login.Username, login.Password)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "invalid credentials")
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
