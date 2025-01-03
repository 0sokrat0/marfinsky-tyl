package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
