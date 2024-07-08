package oauht2

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitOauth2Config() {
	googleProvider := google.New(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("CALLBACK_URL"))
	googleProvider.SetPrompt("select_account")
	goth.UseProviders(googleProvider)
}
