package main

import (
	"go-chat/internal/oauht2"
	"go-chat/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Load env failed")
	}

	oauht2.InitOauth2Config()

	server.Init()
}
