package models

// Gorm Model for the User:
//   - `Model`: regular Model.
//   - `Username`: Username of the user, has a limit of 25 characters.
//   - `Email`: User's email, must be unique, has a limit of 320 characters.
//   - `Password`: User's password, has a limit of 64 characters and should be hashed.
type User struct {
	Model
	Username string `gorm:"type:varchar(25)" json:"username"`
	Email    string `gorm:"type:varchar(320);unique" json:"email"`
	Password string `gorm:"type:varchar(64)" json:"password"`
}
