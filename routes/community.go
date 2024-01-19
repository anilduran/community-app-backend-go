package routes

import (
	"net/http"

	"example.com/community-app-backend/db"
	"example.com/community-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCommunities(c *gin.Context) {

	var communities []models.Community

	result := db.DB.Find(&communities)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": communities,
	})

}

func GetCommunityByID(c *gin.Context) {

	id := c.Param("id")

	var community models.Community

	result := db.DB.First(&community, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, community)

}

func CreateCommunity(c *gin.Context) {

	type CreateCommunityInput struct {
		Name        string `form:"name" binding:"required"`
		Description string `form:"description" binding:"required"`
		ImageUrl    string `form:"image_url" binding:"required"`
	}

	var input CreateCommunityInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	community := models.Community{Name: input.Name, CreatorID: userId}

	result := db.DB.Create(&community)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, community)

}

func UpdateCommunity(c *gin.Context) {

	type UpdateCommunityInput struct {
		Name        string `form:"name"`
		Description string `form:"description"`
		ImageUrl    string `form:"image_url"`
	}

	var input UpdateCommunityInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id := c.Param("id")

	var community models.Community

	result := db.DB.First(&community, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if community.Name != "" {
		community.Name = input.Name
	}

	if community.Description != "" {
		community.Description = input.Description
	}

	if community.ImageUrl != "" {
		community.ImageUrl = input.ImageUrl
	}

	result = db.DB.Save(&community)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, community)

}

func DeleteCommunity(c *gin.Context) {

	id := c.Param("id")

	var community models.Community

	result := db.DB.First(&community, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&community)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
