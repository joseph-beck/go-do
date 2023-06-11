package database

import (
	"fmt"
	"go-do/pkg/util"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StoreReader interface {
	Read(i interface{})
}

type StoreAdder interface {
	Add(i interface{})
}

type StoreUpdater interface {
	Update(i interface{})
}

type StoreDeleter interface {
	Delete(i interface{})
}

type StoreChecker interface {
	Check(i interface{}) bool
}

type StorePinger interface {
	Ping() error
}

type StoreCreator interface {
	Create() error
}

type StoreDestroyer interface {
	Destroy() error
}

type Storer interface {
	StoreAdder
	StoreUpdater
	StoreDeleter
	StoreChecker
	StoreCreator
	StoreDestroyer
}

type Store struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

func MakeStore() *Store {
	log.Println("Making Store")

	return &Store{
		db: makeDb(),
	}
}

func (s *Store) Close() {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	log.Println("Closing Store")

	db, err := s.db.DB()
	util.ErrOut(err)
	err = db.Close()
	util.ErrOut(err)
}

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
