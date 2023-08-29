package models

type User struct {
	Model
	Username string `gorm:"type:varchar(25)" json:"username"`
	Email    string `gorm:"type:varchar(320);unique" json:"email"`
	Password string `gorm:"type:varchar(64)" json:"password"`
}
