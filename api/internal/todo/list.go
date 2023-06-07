package todo

type TodoList struct {
	Items []Item
}

func MakeTodoList() *TodoList {
	return &TodoList{}
}

func (t *TodoList) Str() string {
	return ""
}

func (t *TodoList) Add(item Item) {
	t.Items = append(t.Items, item)
}

func (t *TodoList) Contains(id int) bool {
	for _, item := range t.Items {
		if item.Id == id {
			return true
		}
	}

	return false
}

func (t *TodoList) Remove(id int) {
	for i, item := range t.Items {
		if item.Id != id {
			continue
		}

		t.Items = append(t.Items[:i], t.Items[i+1:]...)
	}
}
