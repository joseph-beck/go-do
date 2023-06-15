package database

import (
	"context"
	"go-do/pkg/util"
)

func (s *BaseStore) Scan(i *interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Model(&i).First(&i)
	util.ErrOut(r.Error)
}

func (s *BaseStore) Read(i *[]interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Find(&i)
	util.ErrOut(r.Error)
}

func (s *BaseStore) Add(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Create(&i)
	util.ErrOut(r.Error)
}

func (s *BaseStore) Update(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *BaseStore) Delete(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *BaseStore) Check(i interface{}, table string) bool {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	return i != nil
}

func (s *BaseStore) Create(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *BaseStore) Destroy(table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

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
