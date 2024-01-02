package routes

import (
	"net/http"

	"example.com/community-app-backend/db"
	"example.com/community-app-backend/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {

	var posts []models.Post

	result := db.DB.Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func CreatePost(c *gin.Context) {

	type CreatePostInput struct {
		Title       string `form:"title" binding:"required"`
		Content     string `form:"content" binding:"required"`
		ImageUrl    string `form:"image_url" binding:"required"`
		CommunityID uint   `form:"community_id" binding:"required"`
	}

	var input CreatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId := c.GetUint("userId")

	post := models.Post{
		Title:       input.Title,
		Content:     input.Content,
		ImageUrl:    input.ImageUrl,
		AuthorID:    userId,
		CommunityID: input.CommunityID,
	}

	result := db.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, post)

}

func UpdatePost(c *gin.Context) {

	type UpdatePostInput struct {
		Title    string `form:"title"`
		Content  string `form:"content"`
		ImageUrl string `form:"image_url"`
	}

	var input UpdatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if post.Title != "" {
		post.Title = input.Title
	}

	if post.Content != "" {
		post.Content = input.Content
	}

	if post.ImageUrl != "" {
		post.ImageUrl = input.ImageUrl
	}

	result = db.DB.Save(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
