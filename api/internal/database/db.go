package database

import (
	"fmt"
	"go-do/pkg/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	UsersTable = "users"
)

// Function that returns pointer to a gorm db
// api can work with various dbs.
type DbMaker func() *gorm.DB

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the dsn.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - DB_ADDR : stores the address (IP)
//   - DB_PORT : stores the port
//   - DB_USER : stores the username
//   - DB_PASS : stores the password
//   - DB_NAME : stores the database name
func PostgresDb() DbMaker {
	return func() *gorm.DB {
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
}

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the db file location.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - SQLITE_DB : stores the location of .db file.
func SQLiteDb() DbMaker {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		util.ErrOut(err)
		return db
	}
}

// Creates an empty Mock database, running any functions on this will cause errors
func MockDb() DbMaker {
	return func() *gorm.DB {
		return &gorm.DB{}
	}
}
