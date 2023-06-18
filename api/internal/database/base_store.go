package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

// BaseStore stores a pointer to a gorm database and a mutex.
//
// The mutex is used to prevent data races.
//   - db : pointer to gorm DB.
//   - dbMu : db mutex.
type BaseStore struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

// Returns a pointer of a BaseStore.
//
// The BaseStore made has a db that uses method makeDb.
func MakeBaseStore() *BaseStore {
	log.Println("Making Base Store")

	return &BaseStore{
		db: makeDb(),
	}
}

// Closes the database connection of the current db within the BaseStore.
func (b *BaseStore) Close() {
	b.dbMu.Lock()
	defer b.dbMu.Unlock()

	log.Println("Closing Base Store")

	db, err := b.db.DB()
	util.ErrOut(err)
	err = db.Close()
	util.ErrOut(err)
}
