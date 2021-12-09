package models

import "time"

type User struct {
	ID          int       `gorm:"primary_key; NOT NULL AUTO_INCREMENT" json:"id"`
	DisplayName string    `json:"display_name"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
