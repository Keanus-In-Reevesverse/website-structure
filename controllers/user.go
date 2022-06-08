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
	//Create request
	var userCreate models.UserRequest

	//Bind request
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	//Request Validator
	if err := models.UserRequestValidator(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	//Password encode && create
	userCreate.Password = services.SHA256Encoder(userCreate.Password)

	//Parse Request to User model
	user := userRequestParse(&userCreate)

	//Create User on DB
	userCreated := database.UserOps(&user, "create")

	//Parse Return
	userReturn := userResponseParse(userCreated)

	//Response Validator
	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	//Return Error
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

	//Password encode && verify
	password := services.SHA256Encoder(user.Password)

	//Gets user by email and password
	database.DB.Table("USER").Select(&user).Where("email = ? AND password = ?", user.Email, password)

	//Edit user on DB
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

	database.DB.Table("USER").Select(&user).Where("email = ? AND password = ?", user.Email, password)

	//Delete user owned by string
	userCreated := database.UserOps(&user, "delete")

	userReturn := userResponseParse(userCreated)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}
