package database

import (
	"context"
	"fmt"
	"go-do/internal/models"
	"go-do/pkg/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	UsersTable = "users"
)

// Function that returns pointer to a gorm db
// api can work with various dbs.
type DbMaker func() *gorm.DB

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the dsn.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - DB_ADDR : stores the address (IP)
//   - DB_PORT : stores the port
//   - DB_USER : stores the username
//   - DB_PASS : stores the password
//   - DB_NAME : stores the database name
func PostgresDb() DbMaker {
	return func() *gorm.DB {
		dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
			os.Getenv("DB_ADDR"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		util.ErrOut(err)
		return db
	}
}

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the db file location.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - SQLITE_DB : stores the location of .db file.
func SQLiteDb() DbMaker {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")))
		util.ErrOut(err)
		return db
	}
}

// Store data, uses gorm db.
type Store struct {
	db *gorm.DB
}

// Creates a new store with the given db maker.
func New(m DbMaker) *Store {
	return &Store{
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
func (s *Store) ListTask(t string) ([]models.TaskModel, error) {
	m := make([]models.TaskModel, 0)
	r := s.db.Table(t).Not("id = ?", 0).Find(&m)
	return m, r.Error
}

// Gets a task from a table of given id.
func (s *Store) GetTask(t string, i int) (models.TaskModel, error) {
	m := models.TaskModel{Id: i}
	r := s.db.Table(t).Find(&m).First(&m)
	return m, r.Error
}

// Adds a given task to the given table.
func (s *Store) AddTask(t string, m models.TaskModel) error {
	r := s.db.Table(t).Create(&m)
	return r.Error
}

// Updates the given task in the given table.
func (s *Store) UpdateTask(t string, m models.TaskModel) error {
	r := s.db.Table(t).Save(&m)
	return r.Error
}

// Deletes a given task id from a given table.
func (s *Store) DeleteTask(t string, i int) error {
	m := models.TaskModel{Id: i}
	r := s.db.Table(t).Delete(&m)
	return r.Error
}

// Checks that the task id given exists in the given table.
func (s *Store) CheckTask(t string, i int) bool {
	m := models.TaskModel{Id: i}
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

// Creates a new table if one does not exist with that name of the given model.
func (s *Store) CreateTable(t string, m interface{}) error {
	err := s.db.Table(t).AutoMigrate(m)
	return err
}
