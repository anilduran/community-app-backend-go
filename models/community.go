package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Community struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	ImageUrl    string    `json:"image_url"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatorID   uuid.UUID `json:"creator_id"`
	Posts       []Post
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (community *Community) BeforeCreate(tx *gorm.DB) (err error) {
	community.ID = uuid.New()
	return
}
