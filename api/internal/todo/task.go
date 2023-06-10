package todo

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Deadline    Deadline
}

func MakeTask() *Task {
	return &Task{
		Id: 123,
	}
}

func (t *Task) Str() string {
	return ""
}
