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

	if err := models.UserRequestValidator(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
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

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}
