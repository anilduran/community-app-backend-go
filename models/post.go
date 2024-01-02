package models

import "time"

type Post struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ImageUrl    string `json:"image_url"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CommunityID uint   `json:"community_id"`
	AuthorID    uint   `json:"author_id"`
	Comments    []Comment
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
