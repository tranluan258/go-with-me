package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func InitDb() *sqlx.DB {
	dbUrl := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "123456", "go_chat")

	conn, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		log.Errorf("Cannot connect db %s", err.Error())
		panic(1)
	}
	return conn
}
