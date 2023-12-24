package models

type Admin struct {
	Model
	Token string `gorm:"type:varchar(64)" json:"token"`
}
