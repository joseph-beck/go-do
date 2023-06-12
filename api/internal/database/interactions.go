package database

import (
	"context"
	"go-do/pkg/util"
)

func (s *Store) Scan(i *interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Model(&i).First(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Read(i *[]interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Table(table).Find(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Add(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	err := s.Create()
	util.ErrOut(err)

	r := s.db.Create(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Update(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *Store) Delete(i interface{}, table string) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *Store) Check(i interface{}, table string) bool {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	return i != nil
}

func (s *Store) Ping() error {
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
