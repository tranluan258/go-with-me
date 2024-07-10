package oauht2

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func InitOauth2Config() {
	googleProvider := google.New(os.Getenv("GG_CLIENT_ID"), os.Getenv("GG_CLIENT_SECRET"), os.Getenv("GG_CALLBACK_URL"))
	googleProvider.SetPrompt("select_account")

	facebookProvider := facebook.New(os.Getenv("FB_CLIENT_ID"), os.Getenv("FB_CLIENT_SECRET"), os.Getenv("FB_CALLBACK_URL"))
	goth.UseProviders(googleProvider, facebookProvider)
}
