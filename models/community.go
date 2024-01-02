package models

import "time"

type Community struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatorID   uint   `json:"creator_id"`
	Posts       []Post
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
