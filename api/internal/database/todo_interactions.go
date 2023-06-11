package database

import (
	"context"
	"go-do/internal/todo"
	"go-do/pkg/util"
)

func (t *TodoStore) Read(i *todo.TaskModel) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(t.table).Model(&i).First(&i)
	util.ErrOut(r.Error)
}

func (t *TodoStore) ReadAll(l *[]todo.TaskModel) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(t.table).Find(&l)
	util.ErrOut(r.Error)
}

func (t *TodoStore) Add(i todo.TaskModel) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	err := t.Create()
	util.ErrOut(err)

	r := t.db.Create(&i)
	util.ErrOut(r.Error)
}

func (t *TodoStore) Update(i todo.TaskModel) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()
}

func (t *TodoStore) Delete(i todo.TaskModel) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()
}

func (t *TodoStore) Check(i todo.TaskModel) bool {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	return false
}

func (t *TodoStore) Ping() error {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	ctx := context.Background()
	db, err := t.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}
