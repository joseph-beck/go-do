package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

type StoreReader interface {
	Read(interface{})
	ReadAll() []interface{}
}

type StoreAdder interface {
	Add(interface{})
}

type StoreUpdater interface {
	Update(interface{})
}

type StoreDeleter interface {
	Delete(interface{})
}

type StoreChecker interface {
	Check(interface{}) bool
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

type Storer interface {
	StoreAdder
	StoreUpdater
	StoreDeleter
	StoreChecker
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
