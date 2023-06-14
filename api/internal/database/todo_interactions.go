package database

import (
	"context"
	"go-do/internal/todo"
	"go-do/pkg/util"
)

// Scans the given task model in from the database
func (t *TodoStore) Scan(task *todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Model(&task).First(&task)
	util.ErrOut(r.Error)
}

// Reads all of the tasks into a given slice of task models.
func (t *TodoStore) Read(tasks *[]todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Find(&tasks)
	util.ErrOut(r.Error)
}

// Adds a task to a given table, if it exists it updates.
// If table doesn't exist, creates that table.
func (t *TodoStore) Add(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Create(&task)
	util.ErrOut(r.Error)

}

// Updates a given task model
func (t *TodoStore) Update(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Save(&task)
	util.ErrOut(r.Error)
}

// Deletes a given task model
func (t *TodoStore) Delete(task todo.TaskModel, table string) {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	r := t.db.Table(table).Delete(&task)
	util.ErrOut(r.Error)
}

// Used to check whether a given task exists.
//
// Can also be used to check whether a table exists,
// pass `todo.TaskModel{Id: 0}` as task.
func (t *TodoStore) Check(task todo.TaskModel, table string) bool {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	e := t.db.Migrator().HasTable(table)
	if task.Id == 0 || !e {
		return e
	}

	r := t.db.Table(table).Model(&task).First(&task)
	return r.Error == nil
}

// Creates a table, with the given model as the schema for the table,
// name of table is given in params
func (t *TodoStore) Create(table string) error {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	err := t.db.Migrator().CreateTable(&todo.TaskModel{})
	util.ErrOut(err)

	task := todo.TaskModel{
		Id: 1,
	}
	r := t.db.Table(table).Create(&task)
	return r.Error
}

// Destroy a given table
func (t *TodoStore) Destroy(table string) error {
	t.dbMu.Lock()
	defer t.dbMu.Unlock()

	return nil
}

// Ping a database
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
