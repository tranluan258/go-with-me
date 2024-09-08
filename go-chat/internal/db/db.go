package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func InitDb() *sqlx.DB {
	driver := os.Getenv("DB_DRIVER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbUserName := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUserName, dbPass, dbName)

	conn, err := sqlx.Connect(driver, dbUrl)
	if err != nil {
		log.Errorf("Cannot connect db %s", err.Error())
		panic(1)
	}
	return conn
}
