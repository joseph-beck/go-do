package database

import "context"

func (s *Store) Add(i interface{}) {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()
}

func (s *Store) Update(i interface{}) {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()
}

func (s *Store) Delete(i interface{}) {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()
}

func (s *Store) Check(i interface{}) bool {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()

	return i != nil
}

func (s *Store) Ping() error {
	s.DbMu.Lock()
	defer s.DbMu.Unlock()

	ctx := context.Background()
	err := s.Db.Ping(ctx)
	return err
}
