package todo

import (
	"fmt"
	"strings"
)

type Deadline struct {
	Date Date
	Time Time
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Time struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

func makeDeadline(s string) *Deadline {
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

	return &Deadline{
		Date: date,
		Time: time,
	}
}

func (d *Deadline) Str() string {
	return fmt.Sprintf(
		"%s-%s",
		d.Date.Str(),
		d.Time.Str(),
	)
}

func (d *Date) Str() string {
	return fmt.Sprintf("%s/%s/%s", d.Day, d.Month, d.Year)
}

func (t *Time) Str() string {
	return fmt.Sprintf("%s:%s", t.Hour, t.Minute)
}
