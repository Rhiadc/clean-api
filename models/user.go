package models

type User struct {
	ID          int    `gorm:"primary_key"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Image       string `json:"image"`
}
