package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

// TodoStore stores a pointer to a gorm database and a mutex.
//
// The mutex is used to prevent data races.
//   - db : pointer to gorm DB.
//   - dbMu : db mutex.
type TodoStore struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

// Returns a pointer of a TodoStore.
//
// The TodoStore made has a db that uses method makeDb.
func MakeTodoStore() *TodoStore {
	log.Println("Making Todo Store")

	return &TodoStore{
		db: makeDb(),
	}
}

// Closes the database connection of the current db within the TodoStore.
func (t *TodoStore) Close() {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	log.Println("Closing Todo Store")

	db, err := t.db.DB()
	util.ErrOut(err)
	err = db.Close()
	util.ErrOut(err)
}
