package database

import "go-do/internal/models"

const (
	UserTable = "users"
)

// Scans a user from a store.
type UserLister[T models.UserModel] interface {
	ListUser(string) ([]T, error)
}

// Reads the entire store into a slice.
type UserGetter[T models.UserModel] interface {
	Read(string, int) (T, error)
}

// Adds a user to a store.
type UserAdder[T models.UserModel] interface {
	AddUser(T) error
}

// Updates a user in a store.
type UserUpdater[T models.UserModel] interface {
	UpdateUser(string, T) error
}

// Deletes a user from a store.
type UserDeleter interface {
	Delete(string, int) error
}

// Checks a user is in a store.
type UserChecker interface {
	Check(string, int) bool
}

// Interface for a Store:
//   - lister : gets a list of Users
//   - getter : gets a User
//   - Adder : adds a user.
//   - Updater : updates a user.
//   - Deleter : deletes a user.
//   - Checker : checks a user exists.
type UserStorer[T models.UserModel] interface {
	UserLister[T]
	UserGetter[T]
	UserAdder[T]
	UserUpdater[T]
	UserDeleter
	UserChecker
}
