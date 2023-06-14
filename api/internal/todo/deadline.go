package todo

import (
	"fmt"
	"go-do/pkg/util"
	"strconv"
	"strings"
)

type Deadline struct {
	Date Date `json:"date"` // date of deadline
	Time Time `json:"time"` // time of deadline
}

func (d *Deadline) Verify() bool {
	return d.Date.Verify() && d.Time.Verify()
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

func (d *Date) Verify() bool {
	w, err := strconv.Atoi(d.Day)
	util.ErrLog(err)
	m, err := strconv.Atoi(d.Month)
	util.ErrLog(err)

	return w <= 31 && m <= 12
}

type Time struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

func (t *Time) Verify() bool {
	h, err := strconv.Atoi(t.Hour)
	util.ErrLog(err)
	m, err := strconv.Atoi(t.Minute)
	util.ErrLog(err)

	return h <= 24 && m <= 60
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
