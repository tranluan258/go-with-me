package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
)

func InitDb() *pgx.Conn {
	dbUrl := "postgres://postgres:123456@localhost:5432/go_chat"

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Errorf("Cannot connect db %s", err.Error())
		panic(1)
	}
	return conn
}
