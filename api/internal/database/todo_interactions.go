package database

import (
	"context"
	"go-do/internal/todo"
	"go-do/pkg/util"
)

func (t *TodoStore) Read(task *todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Model(&task).First(&task)
	util.ErrOut(r.Error)
}

func (t *TodoStore) ReadAll(tasks *[]todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Find(&tasks)
	util.ErrOut(r.Error)
}

func (t *TodoStore) Add(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	err := t.Create()
	util.ErrOut(err)

	r := t.db.Table(table).Create(&task)
	util.ErrOut(r.Error)
}

func (t *TodoStore) Update(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()
}

func (t *TodoStore) Delete(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()
}

func (t *TodoStore) Check(task todo.TaskModel, table string) bool {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Model(&task).First(&task)
	return r.Error == nil
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
