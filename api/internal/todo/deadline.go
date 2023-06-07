package todo

type Deadline struct {
	Date Date
	Time Time
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Time struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func (d *Deadline) Str() string {
	return ""
}

func (d *Date) Str() string {
	return ""
}

func (t *Time) Str() string {
	return ""
}