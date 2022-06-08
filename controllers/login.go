package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/Keanus-In-Reevesverse/website-structure/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no login": err.Error()})
		return
	}

	if err := models.LoginValidator(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var user models.User

	database.DB.Table("USER").Where("email = ?", login.Email).First(&user)

	if user.Email != login.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no banco": "Usuário não existe!"})
		return
	}

	//Encode pass ennter && verify
	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no login": "Credenciais invalidas!"})
		return
	}

	//Generate JWT Token
	token, err := services.NewJWTService().GenerateToken(user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro no login": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token})
}
