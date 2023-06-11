package database

import (
	"go-do/internal/todo"
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

type TodoStoreReader interface {
	Read(*todo.TaskModel)
	ReadAll(*[]todo.TaskModel)
}

type TodoStoreAdder interface {
	Add(todo.TaskModel)
}

type TodoStoreUpdater interface {
	Update(todo.TaskModel)
}

type TodoStoreDeleter interface {
	Delete(todo.TaskModel)
}

type TodoStoreChecker interface {
	Check(todo.TaskModel) bool
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
	TodoStoreAdder
	TodoStoreUpdater
	TodoStoreDeleter
	TodoStoreChecker
	TodoStoreCreator
	TodoStoreDestroyer
	TodoStoreCloser
}

type TodoStore struct {
	db    *gorm.DB
	dbMu  sync.Mutex
	table string
}

func MakeTodoStore(t string) *TodoStore {
	return &TodoStore{
		db:    makeDb(),
		table: t,
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
