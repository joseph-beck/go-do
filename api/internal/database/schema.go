package database

func (s *Store) Create() error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	return nil
}

func (s *Store) Destroy() error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	return nil
}
