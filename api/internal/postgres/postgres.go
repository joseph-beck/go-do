package postgres

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func MakeDb() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	})
}
