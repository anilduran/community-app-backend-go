package db

import (
	"os"

	"example.com/community-app-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (err error) {

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.User{}, &models.Community{}, &models.Post{}, &models.Comment{}, &models.Category{}, &models.Role{})

	return
}
