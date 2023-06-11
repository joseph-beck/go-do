package database

import (
	"go-do/internal/todo"
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

type TodoStoreScanner interface {
	Scan(*todo.TaskModel, string)
}

type TodoStoreReader interface {
	Read(*[]todo.TaskModel, string)
}

type TodoStoreAdder interface {
	Add(todo.TaskModel, string)
}

type TodoStoreUpdater interface {
	Update(todo.TaskModel, string)
}

type TodoStoreDeleter interface {
	Delete(todo.TaskModel, string)
}

type TodoStoreChecker interface {
	Check(todo.TaskModel, string) bool
}

type TodoStorePinger interface {
	Ping() error
}

type TodoStoreCreator interface {
	Create() error
}

type TodoStoreDestroyer interface {
	Destroy() error
}

type TodoStoreCloser interface {
	Close()
}

type TodoStorer interface {
	TodoStoreScanner
	TodoStoreReader
	TodoStoreAdder
	TodoStoreUpdater
	TodoStoreDeleter
	TodoStoreChecker
	TodoStoreCreator
	TodoStoreDestroyer
	TodoStoreCloser
}

type TodoStore struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

func MakeTodoStore() *TodoStore {
	return &TodoStore{
		db: makeDb(),
	}
}

func (t *TodoStore) Close() {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	log.Println("Closing Store")

	db, err := t.db.DB()
	util.ErrOut(err)
	err = db.Close()
	util.ErrOut(err)
}
