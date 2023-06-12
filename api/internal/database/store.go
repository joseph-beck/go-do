package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

type StoreScanner[T any] interface {
	Scan(*T, string)
}

type StoreReader[T any] interface {
	Read(*[]T, string)
}

type StoreAdder[T any] interface {
	Add(T, string)
}

type StoreUpdater[T any] interface {
	Update(T, string)
}

type StoreDeleter[T any] interface {
	Delete(T, string)
}

type StoreChecker[T any] interface {
	Check(T, string) bool
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

type StoreCloser interface {
	Close()
}

type Storer[T any] interface {
	StoreAdder[T]
	StoreUpdater[T]
	StoreDeleter[T]
	StoreChecker[T]
	StoreCreator
	StoreDestroyer
	StoreCloser
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
