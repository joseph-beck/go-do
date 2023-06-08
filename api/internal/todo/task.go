package todo

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Deadline    Deadline
}

func (t *Task) Str() string {
	return ""
}
