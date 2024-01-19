package routes

import (
	"net/http"

	"example.com/community-app-backend/db"
	"example.com/community-app-backend/models"
	"example.com/community-app-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMyCredentials(c *gin.Context) {

	var user models.User

	userId, _ := uuid.Parse(c.GetString("userId"))

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	type UpdateUserInput struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		user.Password = utils.HashPassword(input.Password)
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyPosts(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var posts []models.Post

	result := db.DB.Where("author_id = ?", userId).Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)

}

func GetMyComments(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var comments []models.Comment

	result := db.DB.Where("user_id = ?", userId).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comments)

}

func GetMyCommunities(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var communities []models.Community

	result := db.DB.Where("creator_id = ?", userId).Find(&communities)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, communities)

}
