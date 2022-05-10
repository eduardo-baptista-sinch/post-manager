package domain

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
