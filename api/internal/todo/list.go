package todo

type TodoList struct {
	Tasks []Task
}

func MakeTodoList() *TodoList {
	return &TodoList{}
}

func (t *TodoList) Str() string {
	return ""
}

func (t *TodoList) Add(item Task) {
	t.Tasks = append(t.Tasks, item)
}

func (t *TodoList) Contains(id int) bool {
	for _, item := range t.Tasks {
		if item.Id == id {
			return true
		}
	}

	return false
}

func (t *TodoList) Remove(id int) {
	for i, item := range t.Tasks {
		if item.Id != id {
			continue
		}

		t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
	}
}
