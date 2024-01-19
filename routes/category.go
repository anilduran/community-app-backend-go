package routes

import (
	"net/http"

	"example.com/community-app-backend/db"
	"example.com/community-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCategories(c *gin.Context) {

	var categories []models.Category

	result := db.DB.Find(&categories)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})

}

func GetCategoryByID(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func CreateCategory(c *gin.Context) {

	type CreateCategoryInput struct {
		Name string `form:"name" binding:"required"`
	}

	var input CreateCategoryInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	category := models.Category{Name: input.Name}

	result := db.DB.Create(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, category)

}

func UpdateCategory(c *gin.Context) {

	type UpdateCategory struct {
		Name string `form:"name" binding:"required"`
	}

	var input UpdateCategory

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if category.Name != "" {
		category.Name = input.Name
	}

	result = db.DB.Save(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return

	}

	c.JSON(http.StatusOK, category)

}

func DeleteCategory(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
