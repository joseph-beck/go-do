package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	task := Task{
		Id:          587612,
		Name:        "Dry",
		Description: "dry something up",
		Complete:    false,
		Deadline: Deadline{
			Date: Date{
				Day:   1,
				Month: 11,
				Year:  2023,
			},
			Time: Time{
				Hour:   10,
				Minute: 30,
			},
		},
	}

	assert.NotNil(t, task)
}
