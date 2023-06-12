package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

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
