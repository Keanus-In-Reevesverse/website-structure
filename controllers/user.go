package controllers

import (
	"fmt"
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Creates
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	database.DB.Table("USER").Create(&user)

	CreateLogin(c, &user)
}

func CreateLogin(c *gin.Context, user *models.User) {
	var login models.Login

	login.UserId = user.UserId
	login.Email = user.Email
	login.Password = user.Name

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		fmt.Printf("Erro: %+v", err)
		return
	}

	database.DB.Table("LOGIN").Create(&login)
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
