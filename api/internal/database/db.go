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

func SQLiteDb() DbMaker {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")))
		util.ErrOut(err)
		return db
	}
}

type Store struct {
	db *gorm.DB
}

func New(m DbMaker) *Store {
	return &Store{
		db: m(),
	}
}

func (s *Store) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	return err
}

func (s *Store) Ping() error {
	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}

func (s *Store) ListTask(t string) ([]models.TaskModel, error) {
	m := make([]models.TaskModel, 0)
	r := s.db.Table(t).Not("id = ?", 0).Find(&m)
	return m, r.Error
}

func (s *Store) GetTask(t string, i int) (models.TaskModel, error) {
	m := models.TaskModel{Id: i}
	r := s.db.Table(t).Find(&m).First(&m)
	return m, r.Error
}

func (s *Store) AddTask(t string, m models.TaskModel) error {
	r := s.db.Table(t).Create(&m)
	return r.Error
}

func (s *Store) UpdateTask(t string, m models.TaskModel) error {
	r := s.db.Table(t).Save(&m)
	return r.Error
}

func (s *Store) DeleteTask(t string, i int) error {
	m := models.TaskModel{Id: i}
	r := s.db.Table(t).Delete(&m)

	return r.Error
}

func (s *Store) CheckTask(t string, i int) bool {
	m := models.TaskModel{Id: i}
	r := s.db.Table(t).Model(&m).First(&m)

	return r.Error == nil
}

func (s *Store) CreateTable(t string, m interface{}) error {
	err := s.db.Table(t).AutoMigrate(m)
	return err
}
