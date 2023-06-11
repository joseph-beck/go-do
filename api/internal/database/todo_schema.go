package database

func (t *TodoStore) Create() error {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	return nil
}

func (t *TodoStore) Destroy() error {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	return nil
}
