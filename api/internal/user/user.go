package user

type UserModel struct {
	Id   int
	Name string
}

func (u *UserModel) ToUserPost() UserPost {
	return UserPost{}
}

func (u *UserModel) ToUser() User {
	return User{}
}

type UserPost struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u *UserPost) ToUserModel() UserModel {
	return UserModel{}
}

func (u *UserPost) ToUser() User {
	return User{}
}

type User struct {
	Id   int
	Name string
}

func (u *User) ToUserModel() UserModel {
	return UserModel{}
}

func (u *User) ToUserPost() UserPost {
	return UserPost{}
}