package todo

// Stores a slice of tasks.
type TodoList struct {
	Tasks []Task
}

// Creates a pointer to an empty todo list.
func MakeTodoList() *TodoList {
	return &TodoList{}
}

// Returns a string representation of the todo list.
func (t *TodoList) Str() string {
	return ""
}

// Adds an item to the todo list.
func (t *TodoList) Add(item Task) {
	t.Tasks = append(t.Tasks, item)
}

// Checks to see if the todo list contains the given item.
func (t *TodoList) Contains(id int) bool {
	for _, item := range t.Tasks {
		if item.Id == id {
			return true
		}
	}

	return false
}

// Removes the given id from the todo list.
func (t *TodoList) Remove(id int) {
	for i, item := range t.Tasks {
		if item.Id != id {
			continue
		}

		t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
	}
}

// Returns the length of the todo list.
func (t *TodoList) Len() int {
	return len(t.Tasks)
}
