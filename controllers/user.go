package controllers

import (
	"errors"
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

func userVerifier(user models.User) (models.User, error) {
	database.DB.Table("USER").Select("user_id").Where("email = ?", &user.Email).Scan(&user.UserId)
	if user.UserId != 0 {
		err := errors.New("email ja cadastrado")
		return models.User{}, err
	}
	return user, nil
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

	user, err := userVerifier(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na criação": err.Error()})
		return
	}

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

	//Password encode && verify
	userCreate.Password = services.SHA256Encoder(userCreate.Password)

	user := userRequestParse(&userCreate)

	//Verifica se email já existe
	user, err := userVerifier(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na criação": err.Error()})
		return
	}

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

	userCreate.Password = services.SHA256Encoder(userCreate.Password)

	user := userRequestParse(&userCreate)

	user, err := userVerifier(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	//Delete user owned by string
	userCreated := database.UserOps(&user, "delete")

	userReturn := userResponseParse(userCreated)

	if err := models.UserResponseValidator(&userReturn); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Erro na response": err.Error()})
	}

	c.JSON(http.StatusOK, &userReturn)
}
