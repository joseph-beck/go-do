package models

type User struct {
	Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
