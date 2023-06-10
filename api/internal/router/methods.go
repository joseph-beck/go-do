package router

type Method int

const (
	Undefined Method = iota // 0
	Get                     // 1
	Post                    // 2
	Patch                   // 3
	Delete                  // 4
)
