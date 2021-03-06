package models

import "time"

type Post struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255; not null;unique" json:"title"`
	Content     string    `gorm:"size:255; not null;" json:"content"`
	UserId      int       `gorm:"not null" json:"user_id"`
	PublishedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"published_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// curl localhost:9090/posts -X POST -d '{"title":"sparzas post","content":"somecontent","user_id":"4"}'
