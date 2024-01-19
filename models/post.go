package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	ImageUrl    string    `json:"image_url"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CommunityID uint      `json:"community_id"`
	AuthorID    uuid.UUID `json:"author_id"`
	Comments    []Comment
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New()
	return
}
