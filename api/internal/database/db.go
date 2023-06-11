package database

import (
	"fmt"
	"go-do/pkg/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func makeDb() *gorm.DB {
	dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.ErrOut(err)

	return db
}
