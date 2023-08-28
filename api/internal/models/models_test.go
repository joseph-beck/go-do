package models

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing the makeDeadline method.
//
// This tests to see if all conversions are occuring correctly.
func TestMakeDeadline(t *testing.T) {
	d := makeDeadline("00/00/0000-00:00")
	assert.NotNil(t, d)

	if d.Date.Str() != "00/00/0000" {
		log.Fatalln("Bad date conversion for deadline")
	}

	if d.Time.Str() != "00:00" {
		log.Fatalln("Bad time conversion for deadline")
	}

	if d.Str() != "00/00/0000-00:00" {
		log.Fatalln("Bad deadline conversion for deadline")
	}
}

// Testing the Verify method of the Deadline struct
//
// Verify ensures that the both date and time are correct,
// this uses Verify method from both date and time structs.
func TestVerifyDeadline(t *testing.T) {

}

// Testing the Verify method of the Date struct
//
// Verify ensures that the date is correct.
func TestVerifyDate(t *testing.T) {

}

// Testing the Verify method of the Time struct
//
// Verify ensures that the both time is correct.
func TestVerifyTime(t *testing.T) {

}

func TestTask(t *testing.T) {
	task := Task{
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
	}

	assert.NotNil(t, task)
}
