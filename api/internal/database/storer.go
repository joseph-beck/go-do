package database

import (
	"fmt"
	"go-do/pkg/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Scans an interface from a store.
type StoreScanner[T any] interface {
	Scan(*T, string)
}

// Reads the entire store into a slice.
type StoreReader[T any] interface {
	Read(*[]T, string)
}

// Adds an interface to a store.
type StoreAdder[T any] interface {
	Add(T, string)
}

// Updates an interface in a store.
type StoreUpdater[T any] interface {
	Update(T, string)
}

// Deletes an interface from a store.
type StoreDeleter[T any] interface {
	Delete(T, string)
}

// Checks an interface is in a store.
type StoreChecker[T any] interface {
	Check(T, string) bool
}

// Pings a store.
type StorePinger interface {
	Ping() error
}

// Creates a table in a store.
type StoreCreator interface {
	Create(string)
}

// Destroys a table in a store.
type StoreDestroyer interface {
	Destroy(string)
}

// Closes a store.
type StoreCloser interface {
	Close()
}

// Interface for a Store:
//   - Adder : adds an interface.
//   - Updater : updates an interface.
//   - Deleter : deletes an interface.
//   - Checker : checks an interface exists.
//   - Creator : creates a table.
//   - Destroyer : destroys a table.
//   - Closer : closes a store.
type Storer[T any] interface {
	StoreScanner[T]
	StoreReader[T]
	StoreAdder[T]
	StoreUpdater[T]
	StoreDeleter[T]
	StoreChecker[T]
	StoreCreator
	StoreDestroyer
	StoreCloser
}

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
