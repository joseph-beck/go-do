package user

type UserModel struct {
	Id   int
	Name string
}

type UserPost struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id   int
	Name string
}
