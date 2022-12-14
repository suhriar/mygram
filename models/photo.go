package models

import (
	"time"
)

type Photo struct {
	ID        uint   `json:"id" gorm:"primaryKey;type:integer"`
	Title     string `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url" validate:"required" gorm:"type:varchar(255);not null"`
	UserID    uint   `json:"user_id"`
	Comments  []Comment
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
