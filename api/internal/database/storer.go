package database

// Storer interface, allowing for reading, writing and deleting of data
type Storer interface {
	Reader
	Getter
	Adder
	Updater
	Deleter
	Checker
}

// Returns all values from a table
type Reader interface {
	Read(any, string) error
}

// Returns only a given value from the table
type Getter interface {
	Get(any, string) error
}

// Inserts a value into a given table
type Adder interface {
	Add(any, string) error
}

// Updates a value in the given table
type Updater interface {
	Update(any, string) error
}

// Deletes a value from the given table
type Deleter interface {
	Delete(any, string) error
}

// Checks a value exists in the given table
type Checker interface {
	Check(any, string) error
}
