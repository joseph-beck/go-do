package models

import "time"

// Gorm model for the Task:
//   - `Model` : regular Model.
//   - `Name` : name of the task, defaults to "Task Name", varchar(50) & not null.
//   - `Description` : description of task, defaults to "Task Description", text & not null.
//   - `Complete` : is task complete? defaults to false, boolean & not null.
//   - `Deadline` : deadline of the task.
type Task struct {
	Model
	Name        string    `gorm:"default:'Task Name';type:varchar(50);not null" json:"name"`        // task name
	Description string    `gorm:"default:'Task Description';type:text;not null" json:"description"` // task description
	Complete    bool      `gorm:"default:false;type:boolean;not null" json:"complete"`              // is it complete?
	Deadline    time.Time `json:"deadline"`                                                         // dd/mm/yyyy-hh:mm
}
