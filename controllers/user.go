package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/Keanus-In-Reevesverse/website-structure/services"
	"github.com/gin-gonic/gin"
)

func userResponseParse(user *models.User) models.UserResponse {
	var userReturn models.UserResponse
	userReturn.Name = user.Name
	userReturn.Email = user.Email
	userReturn.PhoneNumber = user.PhoneNumber

	return userReturn
}

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
	user := database.UserOps(&userCreated, "create")

	//Response
	userReturn := userResponseParse(user)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}

func EditUser(c *gin.Context) {
	var userCreated models.UserRequest

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no bind": err.Error()})
		return
	}

	if err := models.UserRequestValidator(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	password := services.SHA256Encoder(userCreated.Password)
	database.DB.First(&userCreated).Where("email = ? AND password = ?", userCreated.Email, password)

	user := database.UserOps(&userCreated, "edit")

	userReturn := userResponseParse(user)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}

func DeleteUser(c *gin.Context) {
	var userCreated models.UserRequest

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no bind": err.Error()})
		return
	}

	if err := models.UserRequestValidator(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	password := services.SHA256Encoder(userCreated.Password)
	database.DB.First(&userCreated).Where("email = ? AND password = ?", userCreated.Email, password)

	user := database.UserOps(&userCreated, "delete")

	userReturn := userResponseParse(user)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}
