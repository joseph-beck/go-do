package database

import (
	"context"
	"go-do/pkg/util"
)

// Scans the given interface model in from the database
func (s *BaseStore) Scan(i *interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Model(&i).First(&i)
	util.ErrOut(r.Error)
}

// Reads all of the interfaces into a given slice of interface models.
func (s *BaseStore) Read(i *[]interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Find(&i)
	util.ErrOut(r.Error)
}

// Adds a interface model to a given table.
func (s *BaseStore) Add(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Create(&i)
	util.ErrOut(r.Error)
}

// Updates a given interface model.
func (s *BaseStore) Update(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Save(&i)
	util.ErrFat(r.Error)
}

// Deletes a given interface model.
func (s *BaseStore) Delete(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Delete(&i)
	util.ErrFat(r.Error)
}

// Used to check whether a given interface model exists.
func (s *BaseStore) Check(i interface{}, table string) bool {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	e := s.db.Migrator().HasTable(table)
	if !e {
		return e
	}

	r := s.db.Table(table).Model(&i).First(&i)
	return r.Error == nil
}

// Creates a table, with the given model as the schema for the table,
// name of table is given in params
//
// TODO: for generic interface
func (s *BaseStore) Create(table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

// Destroy a given table
func (s *BaseStore) Destroy(table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	err := s.db.Table(table).Migrator().DropTable()
	util.ErrFat(err)
}

// Ping a database
func (s *BaseStore) Ping() error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}
