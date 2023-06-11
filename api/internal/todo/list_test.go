package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoList(t *testing.T) {
	l := MakeTodoList()

	assert.NotNil(t, l)
}

func TestTodoListAdd(t *testing.T) {
	l := MakeTodoList()

	assert.NotNil(t, l)

	l.Add(Task{
		Id:          587612,
		Name:        "Dry",
		Description: "dry something up",
		Complete:    false,
		Deadline: Deadline{
			Date: Date{
				Day:   "1",
				Month: "11",
				Year:  "2023",
			},
			Time: Time{
				Hour:   "10",
				Minute: "30",
			},
		},
	})

	assert.Equal(t, l.Len(), 1)
}

func TestTodoListContains(t *testing.T) {
	l := MakeTodoList()

	assert.NotNil(t, l)

	l.Add(Task{
		Id:          587612,
		Name:        "Dry",
		Description: "dry something up",
		Complete:    false,
		Deadline: Deadline{
			Date: Date{
				Day:   "1",
				Month: "11",
				Year:  "2023",
			},
			Time: Time{
				Hour:   "10",
				Minute: "30",
			},
		},
	})

	e := l.Contains(587612)
	assert.True(t, e)

	e = l.Contains(789123)
	assert.False(t, e)
}

func TestTodoListRemove(t *testing.T) {
	l := MakeTodoList()
	assert.NotNil(t, l)

	l.Add(Task{
		Id:          587612,
		Name:        "Dry",
		Description: "dry something up",
		Complete:    false,
		Deadline: Deadline{
			Date: Date{
				Day:   "1",
				Month: "11",
				Year:  "2023",
			},
			Time: Time{
				Hour:   "10",
				Minute: "30",
			},
		},
	})

	e := l.Contains(587612)
	assert.True(t, e)

	l.Remove(587612)

	e = l.Contains(587612)
	assert.False(t, e)
}

func TestTodoListLen(t *testing.T) {
	l := MakeTodoList()
	assert.NotNil(t, l)

	for i := 0; i < 7; i++ {
		l.Add(Task{})
	}

	assert.Equal(t, l.Len(), 7)

	l = MakeTodoList()

	assert.Equal(t, l.Len(), 0)
}
