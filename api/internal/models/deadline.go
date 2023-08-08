package models

import (
	"fmt"
	"go-do/pkg/util"
	"strconv"
	"strings"
)

// Deadline struct stores a Date and Time.
type Deadline struct {
	Date Date `json:"date"` // date of deadline
	Time Time `json:"time"` // time of deadline
}

// Verifies that both the Date and Time are correct using their respective methods.
func (d *Deadline) Verify() bool {
	return d.Date.Verify() && d.Time.Verify()
}

// Returns a string representation of the Deadline.
func (d *Deadline) Str() string {
	return fmt.Sprintf(
		"%s-%s",
		d.Date.Str(),
		d.Time.Str(),
	)
}

// Date struct:
//   - Day : stores the day as a string, so it can start with a 0.
//   - Month : stores the month as a string, so it can start with a 0.
//   - Year : stores the year as a string in case it needs to start with a 0.
type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

// Verifies the date with a series of comparisons,
// the values of the date must first be converted to integers.
func (d *Date) Verify() bool {
	w, err := strconv.Atoi(d.Day)
	util.ErrLog(err)
	m, err := strconv.Atoi(d.Month)
	util.ErrLog(err)

	return w <= 31 && w > 0 && m <= 12 && m > 0
}

// Returns a string representation of the Date.
func (d *Date) Str() string {
	return fmt.Sprintf("%s/%s/%s", d.Day, d.Month, d.Year)
}

// Time struct:
//   - Hour : stores the hour as a string so that it can start with a 0, 24 hr clock.
//   - Minute : stores the mouth as a string so it can start with a 0.
type Time struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

// Verifies the date with a series of comparisons,
// the values of the time must first be converted to integers.
func (t *Time) Verify() bool {
	h, err := strconv.Atoi(t.Hour)
	util.ErrLog(err)
	m, err := strconv.Atoi(t.Minute)
	util.ErrLog(err)

	return h <= 24 && h >= 0 && m <= 60 && m >= 0
}

// Returns a string representation of the Time.
func (t *Time) Str() string {
	return fmt.Sprintf("%s:%s", t.Hour, t.Minute)
}

// Make deadline returns a Deadline struct based off a single string.
//
// The string is parsed and split based off of certain characters.
func makeDeadline(s string) Deadline {
	a := strings.Split(s, "-")
	d, t := strings.Split(a[0], "/"), strings.Split(a[1], ":")

	date := Date{
		Day:   d[0],
		Month: d[1],
		Year:  d[2],
	}
	time := Time{
		Hour:   t[0],
		Minute: t[1],
	}

	return Deadline{
		Date: date,
		Time: time,
	}
}
