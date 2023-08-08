package database

import "go-do/internal/models"

// Scans a task from a store.
type TaskLister[T models.TaskModel] interface {
	ListTask(t string) ([]T, error)
}

// Reads the entire store into a slice.
type TaskGetter[T models.TaskModel] interface {
	Read(t string, i int) (T, error)
}

// Adds a task to a store.
type TaskAdder[T models.TaskModel] interface {
	AddTask(t string, m T) error
}

// Updates a task in a store.
type TaskUpdater[T models.TaskModel] interface {
	UpdateTask(t string, m T) error
}

// Deletes a task from a store.
type TaskDeleter interface {
	Delete(t string, i int) error
}

// Checks a task is in a store.
type TaskChecker interface {
	Check(t string, i int) bool
}

// Interface for a TaskStorer:
//   - lister : gets a list of Tasks
//   - getter : gets a Task
//   - Adder : adds a task.
//   - Updater : updates a task.
//   - Deleter : deletes a task.
//   - Checker : checks a task exists.
type TaskStorer[T models.TaskModel] interface {
	TaskLister[T]
	TaskGetter[T]
	TaskAdder[T]
	TaskUpdater[T]
	TaskDeleter
	TaskChecker
}
