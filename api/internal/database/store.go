package database

type Storer interface {
	Getter
	Adder
	Updater
	Deleter
	Checker
}

type Getter interface {
	Get(a any) error
}

type Adder interface {
	Add(a any) error
}

type Updater interface {
	Update(a any) error
}

type Deleter interface {
	Delete(a any) error
}

type Checker interface {
	Check(a any) error
}
