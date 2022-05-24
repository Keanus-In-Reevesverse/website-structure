package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Creates
func NewUser(c *gin.Context) {
	var userCreated models.UserInput

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	user := database.CreateUser(&userCreated)
	database.CreateLogin(&userCreated)

	var userReturn models.UserOutput
	userReturn.Name = user.Name
	userReturn.Email = user.Email
	userReturn.PhoneNumber = user.PhoneNumber

	c.JSON(http.StatusOK, &userReturn)
}

//Getters
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Table("USER").Find(&users)
	c.JSON(200, &users)
}

func GetUserById(c *gin.Context) {
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
