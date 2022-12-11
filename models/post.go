package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `gorm:"size:255;not null" json:"title"`
	Description string `gorm:"not null;" json:"description"`
	UserId      uint   `gorm:"not null;" json:"userId"`
}
