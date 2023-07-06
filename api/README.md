# `Go-Do API`

## *`Deps`*

- Gin
- Gorm

## *`Usage`*

1. Start Postgres container with \
   `$ make compose`

2. Run API with \
   `$ make run`

## *`Models`*

```go

type Task struct {
    Id          int      `json:"id"`          // task id
    Name        string   `json:"name"`        // task name
    Description string   `json:"description"` // task description
    Complete    bool     `json:"complete"`    // is task complete
    Deadline    Deadline `json:"deadline"`    // task deadline, more data within
}

type Deadline struct {
    Date Date `json:"date"` // date of deadline
    Time Time `json:"time"` // time of deadline
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
```

## *`Todo`*

- Add App to compose.yaml.
- Tidy up models and API.
- More models?
