package todo

import "fmt"

type TaskModel struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Deadline    string `json:"deadline"`
}

// deadline:"dd/mm/yyyy-hh:mm"

func (t *TaskModel) ToTask() *Task {
	return &Task{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Complete:    t.Complete,
		Deadline:    *makeDeadline(t.Deadline),
	}
}

func (t *TaskModel) Str() string {
	return fmt.Sprintf(
		"%d, %s, %s, %t, %s",
		t.Id,
		t.Name,
		t.Description,
		t.Complete,
		t.Deadline,
	)
}

type Task struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Deadline    Deadline
}

func MakeTask() *Task {
	return &Task{}
}

func (t *Task) ToTaskModel() *TaskModel {
	return &TaskModel{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Complete:    t.Complete,
		Deadline:    t.Deadline.Str(),
	}
}

func (t *Task) Str() string {
	return fmt.Sprintf(
		"%d\n%s\n%s\n%t\n%s\n",
		t.Id,
		t.Name,
		t.Description,
		t.Complete,
		t.Deadline.Str(),
	)
}
