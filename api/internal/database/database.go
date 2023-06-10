package database

type StoreAdder interface {
	Add(i interface{})
}

type StoreUpdater interface {
	Update(i interface{})
}

type StoreDeleter interface {
	Delete(i interface{})
}

type StoreChecker interface {
	Check(i interface{})
}

type Storer interface {
	StoreAdder
	StoreUpdater
	StoreDeleter
	StoreChecker
}

type Store struct {
}

func MakeStore() *Store {
	return &Store{}
}
