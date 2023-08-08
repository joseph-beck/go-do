package database

// Create a list.
type ListCreator interface {
	CreateList(t string) error
}

// Destroy a list.
type ListDestroyer interface {
	DestroyList(t string) error
}

// Check if a list exists
type ListChecker interface {
	CheckList(t string) error
}

// Interface for a ListStorer:
//   - creator : creates a list.
//   - destroyer : destroys a list.
//   - checker : checks a list exists.
type ListStorer interface {
	ListCreator
	ListDestroyer
	ListChecker
}
