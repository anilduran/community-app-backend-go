package routes

import (
	"net/http"

	"example.com/community-app-backend/db"
	"example.com/community-app-backend/models"
	"example.com/community-app-backend/utils"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	type SignIn struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignIn

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	isPasswordCorrect := utils.ComparePasswords(user.Password, input.Password)

	if !isPasswordCorrect {
		c.Status(http.StatusUnauthorized)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func SignUp(c *gin.Context) {

	type SignUpInput struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignUpInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var existingUser models.User

	result := db.DB.Where("email = ? OR username = ?", input.Email, input.Username).First(&existingUser)

	if result.Error == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	hashedPassword := utils.HashPassword(input.Password)

	user = models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}
