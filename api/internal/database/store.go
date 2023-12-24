package database

import (
	"context"
	"go-do/internal/models"
	"go-do/pkg/util"

	"gorm.io/gorm"
)

// Store data, uses gorm db.
type Store struct {
	db *gorm.DB
}

// Creates a new store with the given db maker.
func New(m DbMaker) Store {
	return Store{
		db: m(),
	}
}

// Closes the store and returns an error if one occurred.
func (s *Store) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	return err
}

// Pings the db and returns any errors.
func (s *Store) Ping() error {
	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}

// Lists all tasks from a given table./
// It should not return any with id of 0.
func (s *Store) ListTask(t string) ([]models.Task, error) {
	if !s.CheckTable(t) {
		return nil, util.ErrTableDoesNotExist
	}

	m := make([]models.Task, 0)
	r := s.db.Table(t).Not("id = ?", 0).Find(&m)
	return m, r.Error
}

// Gets a task from a table of given id.
func (s *Store) GetTask(t string, i uint) (models.Task, error) {
	m := models.Task{Model: models.Model{ID: i}}
	r := s.db.Table(t).Find(&m).First(&m)
	return m, r.Error
}

// Adds a given task to the given table.
func (s *Store) AddTask(t string, m models.Task) error {
	if !s.CheckTable(t) {
		return util.ErrTableDoesNotExist
	}

	r := s.db.Table(t).Create(&m)
	return r.Error
}

// Updates the given task in the given table.
func (s *Store) UpdateTask(t string, m models.Task) error {
	if !s.CheckTable(t) {
		return util.ErrTableDoesNotExist
	}

	r := s.db.Table(t).Save(&m)
	return r.Error
}

// Deletes a given task id from a given table.
func (s *Store) DeleteTask(t string, i uint) error {
	if !s.CheckTable(t) {
		return util.ErrTableDoesNotExist
	}

	m := models.Task{Model: models.Model{ID: i}}
	r := s.db.Table(t).Delete(&m)
	return r.Error
}

// Checks that the task id given exists in the given table.
func (s *Store) CheckTask(t string, i uint) bool {
	m := models.Task{Model: models.Model{ID: i}}
	r := s.db.Table(t).Model(&m).First(&m)
	return r.Error == nil
}

// Lists all users from the users table.
func (s *Store) ListUser() ([]models.User, error) {
	m := make([]models.User, 0)
	r := s.db.Table(UsersTable).Not("id = ?", 0).Find(&m)
	return m, r.Error
}

// Gets a specific user from the users table.
func (s *Store) GetUser(i uint) (models.User, error) {
	m := models.User{Model: models.Model{ID: i}}
	r := s.db.Table(UsersTable).Find(&m).First(&m)
	return m, r.Error
}

// Adds a user to the users table.
// Password should be hashed before being added.
func (s *Store) AddUser(m models.User) error {
	r := s.db.Table(UsersTable).Create(&m)
	return r.Error
}

// Update a given user
func (s *Store) UpdateUser(m models.User) error {
	r := s.db.Table(UsersTable).Save(&m)
	return r.Error
}

// Deletes the given user from the users table.
func (s *Store) DeleteUser(i uint) error {
	m := models.User{Model: models.Model{ID: i}}
	r := s.db.Table(UsersTable).Delete(&m)
	return r.Error
}

// Checks if the given user exists in the users table.
func (s *Store) CheckUser(i uint) bool {
	m := models.User{Model: models.Model{ID: i}}
	r := s.db.Table(UsersTable).Find(&m)
	return r.Error == nil
}

// Read something into a struct slice from a given table
func (s *Store) Read(i interface{}, t string) error {
	r := s.db.Table(t).Find(i)
	return r.Error
}

// Get something into a struct from a given table
func (s *Store) Get(i interface{}, t string) error {
	r := s.db.Table(t).Find(i).First(i)
	return r.Error
}

// Add something from a struct into a given table
func (s *Store) Add(i interface{}, t string) error {
	r := s.db.Table(t).Create(i)
	return r.Error
}

// Update something from a struct into a given table
func (s *Store) Update(i interface{}, t string) error {
	r := s.db.Table(t).Save(i)
	return r.Error
}

// Delete soemthing from a struct in a given table
func (s *Store) Delete(i interface{}, t string) error {
	r := s.db.Table(t).Delete(i)
	return r.Error
}

// Check if a given value exists in the table
func (s *Store) Check(i interface{}, t string) bool {
	r := s.db.Table(t).Model(i).First(i)
	return r.Error == nil
}

// Creates a new table if one does not exist with that name of the given model.
func (s *Store) CreateTable(t string, m interface{}) error {
	if s.CheckTable(t) {
		return util.ErrTableAlreadyExists
	}

	err := s.db.Table(t).AutoMigrate(m)
	return err
}

// Deletes a given table name
func (s *Store) DeleteTable(t string) error {
	if !s.CheckTable(t) {
		return util.ErrTableDoesNotExist
	}

	err := s.db.Migrator().DropTable(t)
	return err
}

// Checks if the given table name exists in the database
func (s *Store) CheckTable(t string) bool {
	res := s.db.Migrator().HasTable(t)
	return res
}

// Auto migrate all required tables
func (s *Store) AutoMigrate() error {
	err := s.db.AutoMigrate(&models.Admin{})
	if err != nil {
		return err
	}

	err = s.db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
