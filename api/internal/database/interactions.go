package database

import (
	"context"
	"go-do/pkg/util"
)

func (s *Store) Scan(i interface{}) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Model(&i).First(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Read(i []interface{}) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	r := s.db.Find(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Add(i interface{}) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	err := s.Create()
	util.ErrOut(err)

	r := s.db.Create(&i)
	util.ErrOut(r.Error)
}

func (s *Store) Update(i interface{}) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *Store) Delete(i interface{}) {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()
}

func (s *Store) Check(i interface{}) bool {
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
