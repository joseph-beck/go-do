package database

import (
	"go-do/pkg/util"
	"log"
	"sync"

	"gorm.io/gorm"
)

type BaseStore struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

func MakeBaseStore() *BaseStore {
	log.Println("Making Base Store")

	return &BaseStore{
		db: makeDb(),
	}
}

func (b *BaseStore) Close() {
	b.dbMu.Lock()
	defer b.dbMu.Unlock()

	log.Println("Closing Base Store")

	db, err := b.db.DB()
	util.ErrOut(err)
	err = db.Close()
	util.ErrOut(err)
}
