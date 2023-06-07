package todo

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Deadline    Deadline
}

func (i *Item) Str() string {
	return ""
}
