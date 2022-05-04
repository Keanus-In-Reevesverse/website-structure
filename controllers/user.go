package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Creates
func NewUser(c *gin.Context) {
	var userCreated models.UserCreated

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	user := database.CreateUser(&userCreated)
	//database.CreateLogin(&userCreated)

	c.JSON(http.StatusOK, &user)
}

//Getters
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Table("USER").Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("user_id")

	database.DB.Table("USER").First(&user, id)

	if user.UserId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetLogins(c *gin.Context) {
	var login []models.Login
	database.DB.Table("LOGIN").Find(&login)
	c.JSON(200, &login)
}

func GetLogin(c *gin.Context) {
	var login models.Login
	id := c.Params.ByName("user_id")

	database.DB.Table("LOGIN").First(&login, id)

	if login.UserId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, login)
}
