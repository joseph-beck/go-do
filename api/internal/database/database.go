package database

import (
	"go-do/internal/postgres"
	"go-do/pkg/util"
	"log"
	"sync"

	"github.com/go-pg/pg/v10"
)

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

type Storer interface {
	StoreAdder
	StoreUpdater
	StoreDeleter
	StoreChecker
}

type Store struct {
	Db   *pg.DB
	DbMu sync.Mutex
}

func MakeStore() *Store {
	log.Println("Making Store")

	return &Store{
		Db: postgres.MakeDb(),
	}
}

func (s *Store) Close() {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()

	log.Println("Closing Store")

	err := s.Db.Close()
	util.ErrOut(err)
}
