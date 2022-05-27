package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/Keanus-In-Reevesverse/website-structure/services"
	"github.com/gin-gonic/gin"
)

func userRequestParse(userCreated *models.UserRequest) models.User {
	var user models.User
	user.Name = userCreated.Name
	user.Email = userCreated.Email
	user.PhoneNumber = userCreated.PhoneNumber
	user.Password = userCreated.Password

	return user
}

func userResponseParse(user *models.User) models.UserResponse {
	var userReturn models.UserResponse
	userReturn.Name = user.Name
	userReturn.Email = user.Email
	userReturn.PhoneNumber = user.PhoneNumber

	return userReturn
}

//Creates
func NewUser(c *gin.Context) {
	var userCreate models.UserRequest

	//Bind
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	if err := models.UserRequestValidator(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	//Password encode && create
	userCreate.Password = services.SHA256Encoder(userCreate.Password)

	user := userRequestParse(&userCreate)

	userCreated := database.UserOps(&user, "create")

	//Response
	userReturn := userResponseParse(userCreated)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}

func EditUser(c *gin.Context) {
	var userCreate models.UserRequest

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no bind": err.Error()})
		return
	}

	if err := models.UserRequestValidator(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	user := userRequestParse(&userCreate)
	password := services.SHA256Encoder(user.Password)

	database.DB.Table("USER").First(&user).Where("email = ? AND password = ?", user.Email, password)
	userCreated := database.UserOps(&user, "edit")

	userReturn := userResponseParse(userCreated)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}

func DeleteUser(c *gin.Context) {
	var userCreate models.UserRequest

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no bind": err.Error()})
		return
	}

	if err := models.UserRequestValidator(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	user := userRequestParse(&userCreate)
	password := services.SHA256Encoder(user.Password)

	database.DB.Table("USER").First(&user).Where("email = ? AND password = ?", user.Email, password)
	userCreated := database.UserOps(&user, "delete")

	userReturn := userResponseParse(userCreated)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}
