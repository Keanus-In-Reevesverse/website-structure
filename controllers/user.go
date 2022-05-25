package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/Keanus-In-Reevesverse/website-structure/services"
	"github.com/gin-gonic/gin"
)

//Creates
func NewUser(c *gin.Context) {
	var userCreated models.UserRequest

	//Bind
	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	//Password encode && create
	userCreated.Password = services.SHA256Encoder(userCreated.Password)
	user := database.CreateUser(&userCreated)

	//Response
	var userReturn models.UserResponse
	userReturn.Name = user.Name
	userReturn.Email = user.Email
	userReturn.PhoneNumber = user.PhoneNumber

	c.JSON(http.StatusOK, &userReturn)
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
