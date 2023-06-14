package todo

import "fmt"

// Gorm model for the Task:
//
//   - `Id` : id of task, primary key & auto increment.
//   - `Name` : name of the task, defaults to "Task Name", varchar(50) & not null.
//   - `Description` : description of task, defaults to "Task Description", text & not null.
//   - `Complete` : is task complete? defaults to false, boolean & not null.
//   - `Deadline` : deadline of the task.
type TaskModel struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`                             // primary key
	Name        string `gorm:"default:'Task Name';type:varchar(50);not null"`        // task name
	Description string `gorm:"default:'Task Description';type:text;not null"`         // task description
	Complete    bool   `gorm:"default:false;type:boolean;not null"`                  // is it complete?
	Deadline    string `gorm:"default:'00/00/0000-00:00';type:varchar(50);not null"` // dd/mm/yyyy-hh:mm
}

// Makes a pointer to an empty task model.
func MakeTaskModel() *TaskModel {
	return &TaskModel{}
}

// Converts a task model to struct Task.
func (t *TaskModel) ToTask() *Task {
	return &Task{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Complete:    t.Complete,
		Deadline:    *makeDeadline(t.Deadline),
	}
}

// Returns a string that represents this task model.
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

// Task struct:
//
//   - `Id` : id of task.
//   - `Name` : name of the task.
//   - `Description` : description of task.
//   - `Complete` : is task complete?
//   - `Deadline` : deadline of the task.
type Task struct {
	Id          int      `json:"id"`          // task id
	Name        string   `json:"name"`        // task name
	Description string   `json:"description"` // task description
	Complete    bool     `json:"complete"`    // is task complete
	Deadline    Deadline `json:"deadline"`    // task deadline, more data within
}

// Makes a pointer to an empty task
func MakeTask() *Task {
	return &Task{}
}

// Converts a task to struct TaskModel, which is used with gorm.
func (t *Task) ToTaskModel() *TaskModel {
	return &TaskModel{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Complete:    t.Complete,
		Deadline:    t.Deadline.Str(),
	}
}

// Returns a string that represents this task.
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
